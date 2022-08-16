import { ComponentStory, ComponentMeta } from '@storybook/react'
import Spinner from './index'

export default {
  component: Spinner,
  title: 'Loader/Spinner',
} as ComponentMeta<typeof Spinner>

const Template: ComponentStory<typeof Spinner> = (args) => <Spinner {...args} />

export const Primary = Template.bind({})
Primary.args = {}
