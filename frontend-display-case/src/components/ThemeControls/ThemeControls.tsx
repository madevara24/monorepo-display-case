import { FC } from 'react';
import ThemeToggle from '../ThemeToggle/ThemeToggle';
import CyberpunkToggle from '../StyleVariantToggle/StyleVariantToggle';
import styles from './ThemeControls.module.css';

const ThemeControls: FC = () => {
  return (
    <div className={styles.themeControls}>
      <CyberpunkToggle />
      <ThemeToggle />
    </div>
  );
};

export default ThemeControls;
