// @ts-check

import { configure, addDecorator, addParameters } from '@storybook/react'
import { withInfo } from '@storybook/addon-info'
// @ts-ignore
import { themes } from '@storybook/theming'
import { configureActions } from '@storybook/addon-actions'
import '@storybook/addon-console'
// @ts-ignore
import { withConsole } from '@storybook/addon-console'
import { withNotes } from '@storybook/addon-notes'
import { withKnobs } from '@storybook/addon-knobs'

// Load CSS.
import '../../web/src/SourcegraphWebApp.scss'

async function main() {
  // Load all *.story.tsx files.
  require('../src/components/Toggle.story')
  require('../src/components/tokenTextInput/TokenTextInput.story')
  require('../src/components/completion/CompletionWidget.story')
  require('../src/actions/ActionItem.story')
  // if (require.context) {
  //   // In Webpack:
  //   const requireContext = require.context('../src', true, /\.story\.tsx?$/)
  //   for (const storyModule of requireContext.keys()) {
  //     requireContext(storyModule)
  //   }
  // } else {
  //   throw new Error('running storybooks outside of Webpack is not yet supported')
  // }

  // Configure storybooks.
  configure(() => {
    addDecorator(withNotes())
    addDecorator((storyFn, context) => withKnobs(storyFn, context))
    addDecorator((storyFn, context) => withConsole()(storyFn)(context))
    // @ts-ignore
    addParameters({ theme: themes.dark })
    addDecorator(withInfo({ header: false, propTables: false }))

    configureActions({
      depth: 100,
      limit: 20,
    })
  }, module)
}
main().catch(err => console.error(err))
