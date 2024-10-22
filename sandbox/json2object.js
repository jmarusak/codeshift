import fs from 'node:fs';

try {
  const data = fs.readFileSync('codebase.json', 'utf8');
  console.log('JSON string:', data);
  const jsonObject = JSON.parse(data);
  console.log('JavaScript object:', jsonObject);
} catch (err) {
  console.error('Error reading the file:', err);
}
