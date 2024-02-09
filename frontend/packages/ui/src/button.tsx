import { Slot } from "@radix-ui/react-slot";
import { cva, type VariantProps } from "class-variance-authority";
import * as React from "react";
import { sizes } from "../utils/cva";
import { cn } from "../utils/helpers";

const buttonVariants = cva(
  `relative 
  flex items-center justify-center
  cursor-pointer 
  inline-flex 
  items-center 
  space-x-2 
  text-center 
  font-regular 
  ease-out 
  duration-200 
  rounded-md
  outline-none 
  transition-all 
  outline-0 
  focus-visible:outline-4 
  focus-visible:outline-offset-1
  border
  `,
  {
    variants: {
      variant: {
        primary: `bg-primary text-primary-foreground shadow hover:bg-primary/90`,
        secondary: `bg-scale-1200 text-scale-100 hover:text-scale-800 focus-visible:text-scale-600 border-scale-1100 hover:border-scale-900 focus-visible:outline-scale-700 shadow-sm`,
        default: `text-scale-1200 bg-scale-100 hover:bg-scale-300 border-scale-600 hover:border-scale-700 dark:border-scale-700 hover:dark:border-scale-800 dark:bg-scale-500 dark:hover:bg-scale-600 focus-visible:outline-brand-600 shadow-sm`,
        alternative: `text-brand-600 bg-primary text-primary-foreground shadow hover:bg-primary/90-200 hover:bg-primary text-primary-foreground shadow hover:bg-primary/90-400 border-brand-600 hover:border-brand-300 dark:border-brand-400 hover:dark:border-brand-300 focus-visible:border-brand-300 focus-visible:outline-brand-600 shadow-sm`,
        outline: `text-scale-1200 bg-transparent border-scale-600 hover:border-scale-700 dark:border-scale-800 hover:dark:border-scale-900 focus-visible:outline-scale-700`,
        dashed: `text-scale-1200 border border-dashed border-scale-700 hover:border-scale-900 bg-transparent focus-visible:outline-scale-700 shadow-sm`,
        link: `text-brand-600 border border-transparent hover:bg-primary text-primary-foreground shadow hover:bg-primary/90-400 border-opacity-0 bg-opacity-0 dark:bg-opacity-0 shadow-none focus-visible:outline-scale-700`,
        text: `text-scale-1200 hover:bg-scale-500 shadow-none focus-visible:outline-scale-700 border-transparent`,
        danger: `text-red-1100 bg-red-200 border-red-700 hover:border-red-900 hover:bg-red-900 hover:text-lo-contrast focus-visible:outline-red-700 shadow-sm`,
        warning: `text-amber-1100 bg-amber-200 border-amber-700 hover:border-amber-900 hover:bg-amber-900 hover:text-hi-contrast focus-visible:outline-amber-700 shadow-sm`,
      },
      size: {
        ...sizes,
      },
      disabled: {
        true: "opacity-50 cursor-not-allowed pointer-events-none",
      },
    },
    defaultVariants: {
      variant: "primary",
      size: "small",
    },
  },
);

export type ButtonVariantProps = VariantProps<typeof buttonVariants>;
export interface ButtonProps
  extends React.ButtonHTMLAttributes<HTMLButtonElement>,
    Omit<ButtonVariantProps, "disabled"> {
  asChild?: boolean;
}

const Button = React.forwardRef<HTMLButtonElement, ButtonProps>(
  ({ className, variant, size, asChild = false, ...props }, ref) => {
    const Comp = asChild ? Slot : "button";
    return (
      <Comp
        className={cn(buttonVariants({ variant, size, className }))}
        ref={ref}
        {...props}
      />
    );
  },
);
Button.displayName = "Button";

export { Button, buttonVariants };
