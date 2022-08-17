import React from 'react'
import { ComponentStory, ComponentMeta } from '@storybook/react'
import Button from './index'

export default {
  component: Button,
  title: 'Button',
  argTypes: {
    variant: {
      options: ['default', 'outline', 'inline'],
      control: 'radio',
    },
    size: {
      options: ['small', 'medium', 'large'],
      control: 'radio',
    },
  },
} as ComponentMeta<typeof Button>

const Template: ComponentStory<typeof Button> = (args) => (
  <Button
    disabled={args.disabled}
    variant={args.variant}
    size={args.size}
    selected={args.selected}
    loading={args.loading}
  >
    This is a button
  </Button>
)

export const Default = Template.bind({})
Default.args = {
  disabled: false,
  variant: 'default',
  size: 'medium',
  selected: false,
  loading: false,
}
