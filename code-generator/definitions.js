import yaml from 'js-yaml'
import fs from 'fs'

const outputDir = '../definitions'

const doc = yaml.safeLoad(fs.readFileSync('/Users/tyler.liu/src/dotnet/RingCentral.Net/code-generator/rc-platform-adjusted.yml', 'utf8'))
const definitions = doc.definitions
const models = Object.keys(definitions).map(k => ({ name: k, ...definitions[k] })).filter(m => m.name === 'TokenInfo')

for (const model of models) {
  console.log(outputDir, model)
}
