import { createContext, useContext, useState, useEffect, ReactNode } from 'react';

type BaseTheme = 'light' | 'dark';

interface ThemeContextType {
  baseTheme: BaseTheme;
  toggleBaseTheme: () => void;
}

const ThemeContext = createContext<ThemeContextType | undefined>(undefined);

export const ThemeProvider = ({ children }: { children: ReactNode }) => {
  const [baseTheme, setBaseTheme] = useState<BaseTheme>(() => {
    const saved = localStorage.getItem('baseTheme');
    return (saved as BaseTheme) || 'light';
  });

  useEffect(() => {
    localStorage.setItem('baseTheme', baseTheme);
    document.documentElement.classList.toggle('dark-mode', baseTheme === 'dark');
  }, [baseTheme]);

  const toggleBaseTheme = () => {
    setBaseTheme(prev => prev === 'light' ? 'dark' : 'light');
  };

  return (
    <ThemeContext.Provider value={{ 
      baseTheme, 
      toggleBaseTheme
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