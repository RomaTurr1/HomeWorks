import { useState } from 'react';
import './ShoppingList.css';

function ShoppingList() {
  // --- Состояния для добавления ---
  const [title, setTitle] = useState("");
  const [qty, setQty] = useState("");
  const [items, setItems] = useState([]);

  // --- Состояния для фильтрации (Доп. задание) ---
  const [filter, setFilter] = useState("all"); // "all", "bought", "active"

  // --- Состояния для редактирования ---
  const [editId, setEditId] = useState(null);
  const [editTitle, setEditTitle] = useState("");
  const [editQty, setEditQty] = useState("");

  // --- Вычисляемые значения (без useEffect) ---
  const totalCount = items.length;
  const boughtCount = items.filter(item => item.bought).length;
  const remainingCount = totalCount - boughtCount;

  // Логика фильтрации массива перед рендерингом
  const filteredItems = items.filter(item => {
    if (filter === "bought") return item.bought;
    if (filter === "active") return !item.bought;
    return true;
  });

  // --- Функции ---

  const addItem = () => {
    if (title.trim() === "") return;

    const newItem = {
      id: Date.now(),
      title: title.trim(),
      qty: qty || 1, // Если пусто — ставим 1
      bought: false
    };

    setItems([...items, newItem]);
    setTitle("");
    setQty("");
  };

  const deleteItem = (id) => {
    setItems(items.filter(item => item.id !== id));
  };

  const toggleBought = (id) => {
    setItems(items.map(item => 
      item.id === id ? { ...item, bought: !item.bought } : item
    ));
  };

  const startEdit = (item) => {
    setEditId(item.id);
    setEditTitle(item.title);
    setEditQty(item.qty);
  };

  const saveEdit = () => {
    if (editTitle.trim() === "") return;

    setItems(items.map(item => 
      item.id === editId 
        ? { ...item, title: editTitle, qty: editQty || 1 } 
        : item
    ));
    setEditId(null);
  };

  return (
    <div className="shopping">
      <h2>🛒 Shopping List</h2>

      {/* Счётчики (Доп. задание) */}
      <div className="shopping-top">
        <span>Всего: {totalCount}</span>
        <span>Куплено: {boughtCount}</span>
        <span>Осталось: {remainingCount}</span>
      </div>

      {/* Поля ввода нового товара */}
      <div className="shopping-inputs">
        <input 
          className="title" 
          placeholder="Название товара..." 
          value={title}
          onChange={(e) => setTitle(e.target.value)}
        />
        <input 
          className="qty" 
          type="number" 
          placeholder="Кол-во" 
          value={qty}
          onChange={(e) => setQty(e.target.value)}
        />
        <button onClick={addItem}>Добавить</button>
      </div>

      {/* Фильтры (Доп. задание) */}
      <div className="shopping-filters">
        <button 
          className={filter === 'all' ? 'active' : ''} 
          onClick={() => setFilter('all')}
        >Все</button>
        <button 
          className={filter === 'bought' ? 'active' : ''} 
          onClick={() => setFilter('bought')}
        >Куплено</button>
        <button 
          className={filter === 'active' ? 'active' : ''} 
          onClick={() => setFilter('active')}
        >Не куплено</button>
      </div>

      {/* Список товаров */}
      <ul className="shopping-list">
        {filteredItems.map(item => (
          <li key={item.id} className={`shopping-item ${item.bought ? 'done' : ''}`}>
            
            {editId === item.id ? (
              // Режим РЕДАКТИРОВАНИЯ
              <>
                <div className="shopping-inputs" style={{ margin: 0, flex: 1 }}>
                   <input 
                    className="title"
                    value={editTitle} 
                    onChange={(e) => setEditTitle(e.target.value)} 
                   />
                   <input 
                    className="qty"
                    type="number"
                    value={editQty} 
                    onChange={(e) => setEditQty(e.target.value)} 
                   />
                </div>
                <div className="shopping-actions">
                  <button onClick={saveEdit}>💾</button>
                  <button onClick={() => setEditId(null)}>❌</button>
                </div>
              </>
            ) : (
              // Режим ПРОСМОТРА
              <>
                <div className="shopping-left">
                  <input 
                    type="checkbox" 
                    checked={item.bought} 
                    onChange={() => toggleBought(item.id)} 
                  />
                  <span className="shopping-title">{item.title}</span>
                  <span className="shopping-qty">x{item.qty}</span>
                </div>
                <div className="shopping-actions">
                  <button onClick={() => startEdit(item)}>✏️</button>
                  <button onClick={() => deleteItem(item.id)}>🗑️</button>
                </div>
              </>
            )}

          </li>
        ))}
      </ul>
    </div>
  );
}

export default ShoppingList;