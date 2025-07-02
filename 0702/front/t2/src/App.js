import React from 'react';
import { ThemeProvider, useTheme } from './useTheme';
import './App.css';

function ThemeCard() {
  const { theme } = useTheme();
  return (
    <div className={`theme-card ${theme}`} style={{
      padding: '2rem',
      borderRadius: '1rem',
      boxShadow: '0 2px 8px rgba(0,0,0,0.1)',
      background: theme === 'light' ? '#fff' : '#333',
      color: theme === 'light' ? '#222' : '#fff',
      marginBottom: '1rem',
      textAlign: 'center',
      transition: 'all 0.3s',
    }}>
      当前主题：{theme === 'light' ? '● 亮色模式' : '○ 暗色模式'}
    </div>
  );
}

function ThemeButton() {
  const { toggleTheme, theme } = useTheme();
  return (
    <button onClick={toggleTheme} className="theme-btn" style={{
      padding: '0.5rem 1.5rem',
      borderRadius: '2rem',
      border: 'none',
      background: theme === 'light' ? '#222' : '#fff',
      color: theme === 'light' ? '#fff' : '#222',
      cursor: 'pointer',
      fontSize: '1rem',
      transition: 'all 0.3s',
    }}>
      ◑ 切换主题
    </button>
  );
}

function ThemeApp() {
  return (
    <ThemeProvider>
      <div style={{ minHeight: '100vh', display: 'flex', flexDirection: 'column', alignItems: 'center', justifyContent: 'center' }}>
        <ThemeCard />
        <ThemeButton />
      </div>
    </ThemeProvider>
  );
}

export default ThemeApp;
