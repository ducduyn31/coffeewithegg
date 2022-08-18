import { ComponentStory, ComponentMeta } from '@storybook/react'
import Spinner from './index'

export default {
  component: Spinner,
  title: 'Loader/Spinner',
  argTypes: {
    variant: {
      options: ['default', 'button'],
      control: 'radio',
    },
  },
} as ComponentMeta<typeof Spinner>

const Template: ComponentStory<typeof Spinner> = (args) => (
  <Spinner variant={args.variant} {...args} />
)

export const Primary = Template.bind({})
Primary.args = {
  variant: 'default',
}
