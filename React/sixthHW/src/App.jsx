import React, { useReducer } from "react";


function makeId() {
  const time = Date.now().toString(36);
  const rand = Math.floor(Math.random() * 1e6)
    .toString(36)
    .padStart(4, "0");
  return `${time}-${rand}`;
}


const initialState = {
  items: [],
  form: {
    title: "",
    price: "",
  },
  errors: {},
  mode: "add",
  editId: null,
};


function reducer(state, action) {
  switch (action.type) {
    case "SET_FIELD":
      return {
        ...state,
        form: {
          ...state.form,
          [action.field]: action.value,
        },
      };

    case "SET_ERRORS":
      return { ...state, errors: action.errors };

    case "ADD_ITEM":
      return {
        ...state,
        items: [...state.items, action.item],
        form: { title: "", price: "" },
        errors: {},
      };

    case "DELETE_ITEM":
      return {
        ...state,
        items: state.items.filter((it) => it.id !== action.id),
      };

    case "START_EDIT":
      return {
        ...state,
        mode: "edit",
        editId: action.item.id,
        form: {
          title: action.item.title,
          price: action.item.price,
        },
      };

    case "UPDATE_ITEM":
      return {
        ...state,
        items: state.items.map((it) =>
          it.id === state.editId ? action.item : it
        ),
        mode: "add",
        editId: null,
        form: { title: "", price: "" },
        errors: {},
      };

    case "CANCEL_EDIT":
      return {
        ...state,
        mode: "add",
        editId: null,
        form: { title: "", price: "" },
      };

    default:
      return state;
  }
}


export default function App() {
  const [state, dispatch] = useReducer(reducer, initialState);


  function handleSubmit(e) {
    e.preventDefault();

    const errors = {};

    if (!state.form.title.trim()) {
      errors.title = "Введите название";
    }

    if (!state.form.price || Number(state.form.price) <= 0) {
      errors.price = "Введите корректную цену";
    }

    if (Object.keys(errors).length > 0) {
      dispatch({ type: "SET_ERRORS", errors });
      return;
    }

    if (state.mode === "add") {
      dispatch({
        type: "ADD_ITEM",
        item: {
          id: makeId(),
          title: state.form.title,
          price: state.form.price,
        },
      });
    } else {
      dispatch({
        type: "UPDATE_ITEM",
        item: {
          id: state.editId,
          title: state.form.title,
          price: state.form.price,
        },
      });
    }
  }


  const totalPrice = state.items.reduce(
    (sum, it) => sum + Number(it.price),
    0
  );


  return (
    <div style={{ maxWidth: "400px", margin: "20px auto" }}>
      <h2>
        {state.mode === "add"
          ? "Добавить товар"
          : "Редактировать товар"}
      </h2>

      <form onSubmit={handleSubmit}>
        <div>
          <input
            placeholder="Название"
            value={state.form.title}
            onChange={(e) =>
              dispatch({
                type: "SET_FIELD",
                field: "title",
                value: e.target.value,
              })
            }
          />
          {state.errors.title && (
            <div style={{ color: "red" }}>{state.errors.title}</div>
          )}
        </div>

        <div>
          <input
            type="number"
            placeholder="Цена"
            value={state.form.price}
            onChange={(e) =>
              dispatch({
                type: "SET_FIELD",
                field: "price",
                value: e.target.value,
              })
            }
          />
          {state.errors.price && (
            <div style={{ color: "red" }}>{state.errors.price}</div>
          )}
        </div>

        <button type="submit">
          {state.mode === "add" ? "Добавить" : "Сохранить"}
        </button>

        {state.mode === "edit" && (
          <button
            type="button"
            onClick={() => dispatch({ type: "CANCEL_EDIT" })}
          >
            Отмена
          </button>
        )}
      </form>

      <hr />

      {state.items.length === 0 ? (
        <p>Список пуст</p>
      ) : (
        state.items.map((it) => (
          <div
            key={it.id}
            style={{
              border:
                state.mode === "edit" && state.editId === it.id
                  ? "2px solid blue"
                  : "1px solid gray",
              padding: "10px",
              marginTop: "10px",
            }}
          >
            <div>{it.title}</div>
            <div>
              {Number(it.price).toLocaleString("ru-RU")} ₸
            </div>

            <button
              onClick={() =>
                dispatch({ type: "START_EDIT", item: it })
              }
            >
              Редактировать
            </button>

            <button
              onClick={() =>
                dispatch({ type: "DELETE_ITEM", id: it.id })
              }
            >
              Удалить
            </button>
          </div>
        ))
      )}

      <h3>
        Total: {totalPrice.toLocaleString("ru-RU")} ₸
      </h3>
    </div>
  );
}