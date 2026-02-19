function UserCard(props) {
  return (
    <div style={{ border: '1px solid #ccc', margin: '10px', padding: '10px' }}>
      <h3>Имя: {props.name}</h3>
      <p>Возраст: {props.age}</p>
      <p>Город: {props.city}</p>
      <p>Статус: {props.isStudent ? "Студент" : "Не студент"}</p>
    </div>
  );
}
export default UserCard;