import { Button } from '@headlessui/react'
import type { ButtonProps } from '@headlessui/react'
import { forwardRef } from 'react'

type DDButtonProps = ButtonProps & {
  variant?: 'primary' | 'secondary' | 'danger'
  size?: 'sm' | 'md' | 'lg'
}

export const DDButton = forwardRef<HTMLButtonElement, DDButtonProps>(
  (
    { children, className = '', variant = 'primary', size = 'md', ...props },
    ref
  ) => {
    const baseClasses =
      'cursor-pointer rounded font-medium transition-colors box-border border-2 disabled:cursor-auto disabled:bg-gray-300 disabled:text-gray-500 disabled:border-gray-300'

    const variantClasses = {
      primary:
        'bg-darkblue text-white hover:bg-ocean border-darkblue hover:border-ocean',
      secondary:
        'bg-white border-darkblue text-darkblue hover:bg-ocean hover:text-white hover:border-ocean',
      danger:
        'bg-red-600 text-white hover:bg-red-700 border-red-600 hover:border-red-700',
    }

    const sizeClasses = {
      sm: 'px-3 py-1.5 text-sm',
      md: 'px-4 py-2',
      lg: 'px-6 py-3 text-lg',
    }

    const combinedClassName = `${baseClasses} ${variantClasses[variant]} ${sizeClasses[size]} ${className}`

    return (
      <Button ref={ref} className={combinedClassName} {...props}>
        {children}
      </Button>
    )
  }
)

DDButton.displayName = 'DDButton'
