#!/usr/bin/env python3
"""
Documentation Drift Detector
Finds conflicts, duplicates, and outdated information in your Seed
"""

import os
import re
import sys
import json
from datetime import datetime, timedelta
from collections import defaultdict
import hashlib

class DriftDetector:
    def __init__(self, root_path='.'):
        self.root_path = root_path
        self.issues = []
        self.doc_cache = {}
        self.feature_mentions = defaultdict(list)
        self.date_references = []
        self.duplicate_content = []
        
    def scan(self):
        """Main scanning function"""
        print("🔍 Scanning for documentation drift...")
        
        # Find all markdown files
        md_files = self._find_markdown_files()
        print(f"📄 Found {len(md_files)} documentation files")
        
        # Load and analyze each file
        for filepath in md_files:
            self._analyze_file(filepath)
        
        # Cross-file analysis
        self._detect_conflicts()
        self._detect_duplicates()
        self._detect_outdated_dates()
        self._detect_contradictions()
        
        return self.issues
    
    def _find_markdown_files(self):
        """Find all .md files in the tree"""
        md_files = []
        for root, dirs, files in os.walk(self.root_path):
            # Skip hidden directories and common excludes
            dirs[:] = [d for d in dirs if not d.startswith('.') and d not in ['node_modules', 'vendor']]
            
            for file in files:
                if file.endswith('.md'):
                    md_files.append(os.path.join(root, file))
        return md_files
    
    def _analyze_file(self, filepath):
        """Analyze a single file for drift indicators"""
        try:
            with open(filepath, 'r', encoding='utf-8') as f:
                content = f.read()
                self.doc_cache[filepath] = {
                    'content': content,
                    'lines': content.split('\n'),
                    'hash': hashlib.md5(content.encode()).hexdigest()[:8]
                }
                
                # Extract features mentioned
                self._extract_feature_mentions(filepath, content)
                
                # Extract date references
                self._extract_dates(filepath, content)
                
                # Check for drift indicators
                self._check_single_file_issues(filepath, content)
                
        except Exception as e:
            self.issues.append({
                'type': 'error',
                'severity': 'warning',
                'file': filepath,
                'message': f'Could not read file: {e}'
            })
    
    def _extract_feature_mentions(self, filepath, content):
        """Extract mentions of features/systems"""
        # Common feature patterns
        patterns = [
            (r'(?:implemented|built|created|finished)\s+(?:the\s+)?(\w+[\s\w]*)', 'implemented'),
            (r'(?:will|plan to|going to)\s+(?:implement|build|create)\s+(?:the\s+)?(\w+[\s\w]*)', 'planned'),
            (r'(?:TODO|FIXME):\s*(.+)', 'todo'),
            (r'✅\s*(.+)', 'completed'),
            (r'❌\s*(.+)', 'not_implemented'),
            (r'🔧\s*(.+)', 'partial'),
        ]
        
        for pattern, status in patterns:
            for match in re.finditer(pattern, content, re.IGNORECASE):
                feature = match.group(1).strip()
                self.feature_mentions[feature.lower()].append({
                    'file': filepath,
                    'status': status,
                    'context': match.group(0)
                })
    
    def _extract_dates(self, filepath, content):
        """Extract date references"""
        # Various date patterns
        date_patterns = [
            r'\b(\d{4}-\d{2}-\d{2})\b',
            r'\b(January|February|March|April|May|June|July|August|September|October|November|December)\s+\d{1,2},?\s+\d{4}\b',
            r'\b(next week|next month|tomorrow|yesterday)\b',
            r'\b(\d+)\s+(days?|weeks?|months?)\s+ago\b',
        ]
        
        for pattern in date_patterns:
            for match in re.finditer(pattern, content, re.IGNORECASE):
                self.date_references.append({
                    'file': filepath,
                    'date_text': match.group(0),
                    'line': content[:match.start()].count('\n') + 1
                })
    
    def _check_single_file_issues(self, filepath, content):
        """Check for issues within a single file"""
        lines = content.split('\n')
        
        # Check for conflicting status markers
        if '✅' in content and '❌' in content:
            completed = len(re.findall(r'✅', content))
            not_done = len(re.findall(r'❌', content))
            if completed > 0 and not_done > 0:
                # Look for the same feature marked differently
                for i, line in enumerate(lines):
                    if '✅' in line:
                        feature = line.replace('✅', '').strip()
                        for j, other_line in enumerate(lines):
                            if '❌' in other_line and feature in other_line:
                                self.issues.append({
                                    'type': 'contradiction',
                                    'severity': 'high',
                                    'file': filepath,
                                    'message': f'Same feature marked as both completed and not implemented',
                                    'lines': [i+1, j+1]
                                })
        
        # Check for outdated TODOs
        todo_pattern = re.compile(r'TODO.*(\d{4}-\d{2}-\d{2})')
        for i, line in enumerate(lines):
            match = todo_pattern.search(line)
            if match:
                todo_date = datetime.strptime(match.group(1), '%Y-%m-%d')
                if (datetime.now() - todo_date).days > 30:
                    self.issues.append({
                        'type': 'outdated',
                        'severity': 'medium',
                        'file': filepath,
                        'line': i+1,
                        'message': f'TODO from {match.group(1)} is over 30 days old'
                    })
    
    def _detect_conflicts(self):
        """Detect conflicts across files"""
        # Check for features with conflicting statuses
        for feature, mentions in self.feature_mentions.items():
            if len(mentions) > 1:
                statuses = set(m['status'] for m in mentions)
                if len(statuses) > 1 and ('implemented' in statuses and 'planned' in statuses):
                    self.issues.append({
                        'type': 'conflict',
                        'severity': 'high',
                        'feature': feature,
                        'message': f'Feature "{feature}" has conflicting statuses',
                        'occurrences': [{'file': m['file'], 'status': m['status']} for m in mentions]
                    })
    
    def _detect_duplicates(self):
        """Detect duplicate content across files"""
        # Simple duplicate detection - look for identical paragraphs
        paragraphs = defaultdict(list)
        
        for filepath, data in self.doc_cache.items():
            content = data['content']
            # Split into paragraphs
            paras = re.split(r'\n\s*\n', content)
            for para in paras:
                para = para.strip()
                if len(para) > 100:  # Only check substantial paragraphs
                    para_hash = hashlib.md5(para.encode()).hexdigest()
                    paragraphs[para_hash].append({
                        'file': filepath,
                        'content': para[:100] + '...' if len(para) > 100 else para
                    })
        
        # Report duplicates
        for para_hash, occurrences in paragraphs.items():
            if len(occurrences) > 1:
                self.duplicate_content.append({
                    'type': 'duplicate',
                    'severity': 'medium',
                    'message': 'Duplicate content found',
                    'preview': occurrences[0]['content'],
                    'files': [o['file'] for o in occurrences]
                })
    
    def _detect_outdated_dates(self):
        """Detect outdated date references"""
        for date_ref in self.date_references:
            date_text = date_ref['date_text'].lower()
            
            # Check relative dates
            if 'next week' in date_text or 'tomorrow' in date_text:
                # Check file modification time
                file_mtime = os.path.getmtime(date_ref['file'])
                file_age_days = (datetime.now() - datetime.fromtimestamp(file_mtime)).days
                if file_age_days > 7:
                    self.issues.append({
                        'type': 'outdated',
                        'severity': 'medium',
                        'file': date_ref['file'],
                        'line': date_ref['line'],
                        'message': f'Relative date "{date_ref["date_text"]}" in {file_age_days}-day old file'
                    })
            
            # Check absolute dates in the past
            if re.match(r'\d{4}-\d{2}-\d{2}', date_text):
                try:
                    date = datetime.strptime(date_text, '%Y-%m-%d')
                    if date < datetime.now() - timedelta(days=90):
                        self.issues.append({
                            'type': 'outdated',
                            'severity': 'low',
                            'file': date_ref['file'],
                            'line': date_ref['line'],
                            'message': f'Date {date_text} is over 90 days old'
                        })
                except:
                    pass
    
    def _detect_contradictions(self):
        """Detect logical contradictions"""
        # Look for conflicting statements about implementation status
        for filepath, data in self.doc_cache.items():
            content = data['content'].lower()
            
            # Common contradiction patterns
            if 'not implemented' in content and 'fully implemented' in content:
                self.issues.append({
                    'type': 'contradiction',
                    'severity': 'high',
                    'file': filepath,
                    'message': 'File contains both "not implemented" and "fully implemented"'
                })
            
            if 'planned' in content and 'completed' in content:
                # More nuanced check - could be legitimate
                planned_count = content.count('planned')
                completed_count = content.count('completed')
                if planned_count > 2 and completed_count > 2:
                    self.issues.append({
                        'type': 'ambiguity',
                        'severity': 'medium',
                        'file': filepath,
                        'message': f'Mixed status: {planned_count} "planned", {completed_count} "completed"'
                    })
    
    def report(self):
        """Generate a report of findings"""
        if not self.issues and not self.duplicate_content:
            print("\n✅ No documentation drift detected!")
            return
        
        print(f"\n⚠️  Found {len(self.issues) + len(self.duplicate_content)} drift issues:\n")
        
        # Group by severity
        by_severity = defaultdict(list)
        for issue in self.issues + self.duplicate_content:
            by_severity[issue['severity']].append(issue)
        
        # Report by severity
        severity_icons = {'high': '🔴', 'medium': '🟡', 'low': '🔵'}
        
        for severity in ['high', 'medium', 'low']:
            if severity in by_severity:
                print(f"{severity_icons[severity]} {severity.upper()} ({len(by_severity[severity])} issues)")
                for issue in by_severity[severity][:5]:  # Show first 5
                    self._print_issue(issue)
                if len(by_severity[severity]) > 5:
                    print(f"   ... and {len(by_severity[severity]) - 5} more\n")
                print()
    
    def _print_issue(self, issue):
        """Print a single issue"""
        if issue['type'] == 'conflict':
            print(f"   📍 {issue['feature']}: {issue['message']}")
            for occ in issue.get('occurrences', [])[:3]:
                print(f"      - {occ['file']}: {occ['status']}")
        elif issue['type'] == 'duplicate':
            print(f"   📄 {issue['message']}")
            print(f"      Preview: {issue['preview']}")
            for f in issue['files'][:3]:
                print(f"      - {f}")
        else:
            file_part = issue['file'].replace(self.root_path + '/', '')
            line_part = f":{issue.get('line', '')}" if 'line' in issue else ''
            print(f"   📍 {file_part}{line_part}")
            print(f"      {issue['message']}")
    
    def save_report(self, output_file='drift-report.json'):
        """Save detailed report as JSON"""
        report = {
            'scan_date': datetime.now().isoformat(),
            'root_path': self.root_path,
            'files_scanned': len(self.doc_cache),
            'total_issues': len(self.issues) + len(self.duplicate_content),
            'issues': self.issues,
            'duplicates': self.duplicate_content,
            'summary': {
                'high': len([i for i in self.issues if i['severity'] == 'high']),
                'medium': len([i for i in self.issues if i['severity'] == 'medium']),
                'low': len([i for i in self.issues if i['severity'] == 'low'])
            }
        }
        
        with open(output_file, 'w') as f:
            json.dump(report, f, indent=2)
        print(f"\n💾 Detailed report saved to {output_file}")


def main():
    # Parse arguments
    path = sys.argv[1] if len(sys.argv) > 1 else '.'
    
    print(f"🌊 Documentation Drift Detector")
    print(f"📂 Scanning: {os.path.abspath(path)}")
    print("=" * 50)
    
    # Run detection
    detector = DriftDetector(path)
    detector.scan()
    detector.report()
    
    # Optionally save detailed report
    if '--save' in sys.argv:
        detector.save_report()


if __name__ == '__main__':
    main()