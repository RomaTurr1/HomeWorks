import { useState } from 'react';
import styles from './App.module.css';
import ProductForm from './components/ProductForm';
import ProductList from './components/ProductList';

function App() {
  const [products, setProducts] = useState([]);

  // 25-40 мин: Добавление (spread)
  const addProduct = (newProduct) => {
    setProducts(prev => [newProduct, ...prev]);
  };

  // 40-55 мин: Удаление
  const removeProduct = (id) => {
    setProducts(prev => prev.filter(p => p.id !== id));
  };

  // 55-70 мин: Изменение qty
  const updateQty = (id, delta) => {
    setProducts(prev => prev.map(p => {
      if (p.id === id) {
        const newQty = p.qty + delta;
        return { ...p, qty: newQty > 0 ? newQty : 1 }; // Минимум 1
      }
      return p;
    }));
  };

  // 70-80 мин: Total (reduce)
  const total = products.reduce((sum, p) => sum + p.price * p.qty, 0);

  return (
    <div className={styles.container}>
      <h2>🛒 Mini Shop Cart</h2>
      <ProductForm onAdd={addProduct} />
      
      <ProductList 
        products={products} 
        onRemove={removeProduct} 
        onUpdateQty={updateQty} 
      />

      <div className={styles.total}>
        Итого: {total.toLocaleString()} руб.
      </div>
    </div>
  );
}

export default App;