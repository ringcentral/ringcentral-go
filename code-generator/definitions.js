import yaml from 'js-yaml'
import fs from 'fs'
import pascalCase from 'pascal-case'
import toStartCase from 'to-start-case'
import path from 'path'

const outputDir = '../definitions'

const doc = yaml.safeLoad(fs.readFileSync('/Users/tyler.liu/src/dotnet/RingCentral.Net/code-generator/rc-platform-adjusted.yml', 'utf8'))
const definitions = doc.definitions
const models = Object.keys(definitions).map(k => ({ name: k, ...definitions[k] })).filter(m => m.type !== 'array')

const normalizeType = p => {
  if (p.type === 'integer') {
    return 'int'
  } else if (p.type === 'array') {
    return `[]${normalizeType(p.items)}`
  } else if (p.type === undefined || p.type === 'object') {
    if (!p.$ref) {
      return 'interface{}' // anonymous object
    }
    return p.$ref.split('/').slice(-1)[0]
  } else if (p.type === 'boolean') {
    return 'bool'
  } else if (p.type === 'file') {
    return 'Attachment'
  } else if (p.type === 'string') {
    return 'string'
  } else {
    throw new Error(`Unknown type ${p.type}`)
  }
}

for (const model of models) {
  const code = `package definitions

// ${model.name} ${toStartCase(model.name)}
type ${model.name} struct {
${
  Object.keys(model.properties).map(k => ({ name: k, ...model.properties[k] }))
    .map(p => `\t${pascalCase(p.name)} ${normalizeType(p)} \`json:"${p.name}"\``).join('\n')
}
}
`
  fs.writeFileSync(path.join(outputDir, `${model.name.toLowerCase()}.go`), code)
}
