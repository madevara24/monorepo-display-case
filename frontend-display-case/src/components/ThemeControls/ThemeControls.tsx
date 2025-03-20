import { FC } from 'react';
import styles from './ThemeControls.module.css';
import { useTheme } from '../../context/ThemeContext';

const ThemeControls: FC = () => {
  const { baseTheme, toggleBaseTheme } = useTheme();

  return (
    <div className={styles.themeControls}>
      <button
        onClick={toggleBaseTheme}
        aria-label={`Switch to ${baseTheme === 'light' ? 'dark' : 'light'} mode`}
      >
        {baseTheme === 'light' ? '🌙' : '☀️'}
      </button>
    </div>
  );
};

export default ThemeControls;
