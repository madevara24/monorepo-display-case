import { FC } from 'react';
import { useTheme } from '../../context/ThemeContext';
import styles from './StyleVariantToggle.module.css';

const CyberpunkToggle: FC = () => {
  const { styleVariant, toggleCyberpunkTheme } = useTheme();
  const isActive = styleVariant === 'cyberpunk';

  return (
    <button
      className={`${styles.styleToggle} ${isActive ? styles.active : ''}`}
      onClick={toggleCyberpunkTheme}
      aria-label={`${isActive ? 'Return to default theme' : 'Try Cyberpunk theme'}`}
    >
      {isActive ? '↩️' : '🎮'}
    </button>
  );
};

export default CyberpunkToggle;