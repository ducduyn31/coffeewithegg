import { Meta, moduleMetadata, Story } from '@storybook/angular'
import { IconButtonComponent } from './icon-button.component'
import icons from '../icon/icon.list'
import { IconComponent } from '../icon/icon.component'

export default {
  title: 'IconButtons',
  component: IconButtonComponent,
  decorators: [
    moduleMetadata({
      declarations: [IconComponent],
      imports: [],
    }),
  ],
} as Meta<IconButtonComponent>

const Template: Story<IconButtonComponent> = (args: IconButtonComponent) => ({
  props: args,
  template: `
      <div style='display: grid;grid-template-columns: repeat(6, minmax(0, 1fr))'>
        ${icons.map(
          (name) => `
          <div style='display: flex; flex-direction: row; justify-content: center; align-items: center; column-gap: 20px;'>
            <cwe-icon-button name="${name}" variant="${args.variant}"></cwe-icon-button>
            <p>${name}</p>
          </div>
        `,
        )}
      </div>
    `,
})

export const Default = Template.bind({})
Default.args = {
  variant: 'default',
}

export const Transparent = Template.bind({})
Transparent.args = {
  variant: 'transparent',
}
