import { useContext, useState, useCallback, useEffect } from 'react';
import ThemeContext from './ThemeContext';

export function ThemeProvider({ children }) {
  const [theme, setTheme] = useState('light');

  const toggleTheme = useCallback(() => {
    setTheme(prev => (prev === 'light' ? 'dark' : 'light'));
  }, []);

  useEffect(() => {
    document.body.style.background = theme === 'light' ? '#fff' : '#222';
    document.body.style.color = theme === 'light' ? '#222' : '#fff';
  }, [theme]);

  return (
    <ThemeContext.Provider value={{ theme, toggleTheme }}>
      {children}
    </ThemeContext.Provider>
  );
}

export function useTheme() {
  return useContext(ThemeContext);
}
