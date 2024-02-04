import { cva, VariantProps } from "class-variance-authority";
import { HTMLAttributes, PropsWithChildren } from "react";
import styles from "./text.module.css";

const text = cva(styles.text, {
  variants: {
    size: {
      1: styles["text-1"],
      2: styles["text-2"],
      3: styles["text-3"],
      4: styles["text-4"],
      5: styles["text-5"],
      6: styles["text-6"],
      7: styles["text-7"],
      8: styles["text-8"],
      9: styles["text-9"],
      10: styles["text-10"],
    },
    bold: {
      true: styles["text-bold"],
    },
  },
  defaultVariants: {
    size: 2,
  },
});

type TextProps = PropsWithChildren<VariantProps<typeof text>> &
  HTMLAttributes<HTMLSpanElement>;

export function Text({ children, className, size, bold, ...props }: TextProps) {
  return (
    <span className={text({ size, className, bold })} {...props}>
      {children}
    </span>
  );
}
