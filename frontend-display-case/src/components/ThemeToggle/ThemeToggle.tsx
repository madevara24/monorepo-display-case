import { FC } from 'react';
import { useTheme } from '../../context/ThemeContext';
import styles from './ThemeToggle.module.css';

const ThemeToggle: FC = () => {
  const { baseTheme, toggleBaseTheme } = useTheme();

  return (
    <button
      className={styles.themeToggle}
      onClick={toggleBaseTheme}
      aria-label={`Switch to ${baseTheme === 'light' ? 'dark' : 'light'} mode`}
    >
      {baseTheme === 'light' ? '🌙' : '🌞'}
    </button>
  );
};

export default ThemeToggle;