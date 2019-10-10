import yaml from 'js-yaml'
import fs from 'fs'
import pascalCase from 'pascal-case'
import toStartCase from 'to-start-case'
import path from 'path'

const outputDir = '../definitions'

const doc = yaml.safeLoad(fs.readFileSync('/Users/tyler.liu/src/dotnet/RingCentral.Net/code-generator/rc-platform-adjusted.yml', 'utf8'))
const definitions = doc.definitions
const models = Object.keys(definitions).map(k => ({ name: k, ...definitions[k] })).filter(m => m.name === 'TokenInfo')

const normalizeType = type => {
  switch (type) {
    case 'integer':
      return 'int'
    default:
      return type
  }
}

for (const model of models) {
  const code = `package definitions

// ${model.name} ${toStartCase(model.name)}
type ${model.name} struct {
${
  Object.keys(model.properties).map(k => ({ name: k, ...model.properties[k] }))
    .map(p => `\t${pascalCase(p.name)} ${normalizeType(p.type)} \`json:"${p.name}"\``).join('\n')
}
}
`
  fs.writeFileSync(path.join(outputDir, `${model.name.toLowerCase()}.go`), code)
}
