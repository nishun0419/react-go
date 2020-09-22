import React from 'react';

let style = {
  maxWidth: '700px',
};

let btn = {
  cursor: 'pointer'
};

const List = (props) => (
  <ul className="siimple-list">
  	{props.todos.map((todo, i) => {
 		return <li key={todo.ID} className="siimple-list-item siimple--bg-white" style={style}>{todo.Name} <span className="siimple-tag siimple-tag--error siimple-hover" style={btn} onClick={() => props.handleRemove(i,todo.ID)}>Delete</span></li>
 	})}
  </ul>
);

export default List;