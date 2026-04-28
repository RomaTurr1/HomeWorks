import { useState } from "react";
import { useDispatch, useSelector } from "react-redux";
import { addToCart, removeFromCart, clearCart } from "./cartSlice";

const productsData = [
  { id: 1, title: "Телефон", price: 120000, category: "tech" },
  { id: 2, title: "Ноутбук", price: 500000, category: "tech" },
  { id: 3, title: "Футболка", price: 8000, category: "clothes" },
  { id: 4, title: "Кроссовки", price: 30000, category: "clothes" },
];

export default function App() {
  const dispatch = useDispatch();
  const cart = useSelector((state) => state.cart.items);

  const [search, setSearch] = useState("");
  const [filter, setFilter] = useState("all");

  /* фильтр + поиск */
  const filteredProducts = productsData.filter((p) => {
    const matchSearch = p.title.toLowerCase().includes(search.toLowerCase());
    const matchFilter = filter === "all" || p.category === filter;
    return matchSearch && matchFilter;
  });

  const totalCount = cart.reduce((sum, i) => sum + i.quantity, 0);

  return (
    <div style={{ padding: 20 }}>
      <h1>Каталог</h1>

      {/* Поиск */}
      <input
        placeholder="Поиск..."
        value={search}
        onChange={(e) => setSearch(e.target.value)}
      />

      {/* Фильтр */}
      <select onChange={(e) => setFilter(e.target.value)}>
        <option value="all">Все</option>
        <option value="tech">Техника</option>
        <option value="clothes">Одежда</option>
      </select>

      <h2>Товары</h2>

      {filteredProducts.map((p) => (
        <div key={p.id} style={{ border: "1px solid", margin: 10, padding: 10 }}>
          <h3>{p.title}</h3>
          <p>{p.price.toLocaleString()} ₸</p>

          <button onClick={() => dispatch(addToCart(p))}>
            В корзину
          </button>
        </div>
      ))}

      <hr />

      <h2>Корзина ({totalCount})</h2>

      {cart.length === 0 ? (
        <p>Пусто</p>
      ) : (
        <>
          {cart.map((item) => (
            <div key={item.id}>
              {item.title} x {item.quantity}
              <button onClick={() => dispatch(removeFromCart(item.id))}>
                ❌
              </button>
            </div>
          ))}

          <button onClick={() => dispatch(clearCart())}>
            Очистить корзину
          </button>
        </>
      )}
    </div>
  );
}