import { useState, useEffect } from 'react';
import Header from './components/Header/Header';
import ProductsPage from './components/Products/ProductsPage';

function App() {
  // useEffect №2: Тема из localStorage
  const [theme, setTheme] = useState(localStorage.getItem('theme') || 'light');

  useEffect(() => {
    document.documentElement.setAttribute('data-theme', theme);
    localStorage.setItem('theme', theme);
  }, [theme]);

  return (
    <div className="app">
      <Header theme={theme} setTheme={setTheme} />
      <ProductsPage />
    </div>
  );
}

export default App;