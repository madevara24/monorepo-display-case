import { createContext, useContext, useState, useEffect, ReactNode } from 'react';

type BaseTheme = 'light' | 'dark';
type StyleVariant = 'default' | 'cyberpunk';

interface ThemeContextType {
  baseTheme: BaseTheme;
  styleVariant: StyleVariant;
  toggleBaseTheme: () => void;
  setStyleVariant: (variant: StyleVariant) => void;
  toggleCyberpunkTheme: () => void;
}

const ThemeContext = createContext<ThemeContextType | undefined>(undefined);

export const ThemeProvider = ({ children }: { children: ReactNode }) => {
  const [baseTheme, setBaseTheme] = useState<BaseTheme>(() => {
    const saved = localStorage.getItem('baseTheme');
    return (saved as BaseTheme) || 'light';
  });

  const [styleVariant, setStyleVariant] = useState<StyleVariant>(() => {
    const saved = localStorage.getItem('styleVariant');
    return (saved as StyleVariant) || 'default';
  });

  useEffect(() => {
    localStorage.setItem('baseTheme', baseTheme);
    localStorage.setItem('styleVariant', styleVariant);
    
    // Apply theme classes
    document.documentElement.classList.toggle('dark-mode', baseTheme === 'dark');
    document.documentElement.classList.toggle('theme-cyberpunk', styleVariant === 'cyberpunk');
  }, [baseTheme, styleVariant]);

  const toggleBaseTheme = () => {
    setBaseTheme(prev => prev === 'light' ? 'dark' : 'light');
  };

  const toggleCyberpunkTheme = () => {
    setStyleVariant(prev => prev === 'cyberpunk' ? 'default' : 'cyberpunk');
  };

  return (
    <ThemeContext.Provider value={{ 
      baseTheme, 
      styleVariant, 
      toggleBaseTheme, 
      setStyleVariant,
      toggleCyberpunkTheme
    }}>
      {children}
    </ThemeContext.Provider>
  );
};

export const useTheme = () => {
  const context = useContext(ThemeContext);
  if (context === undefined) {
    throw new Error('useTheme must be used within a ThemeProvider');
  }
  return context;
}; 