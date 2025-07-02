import React, { useRef, useState, useEffect, useMemo } from 'react';
import { TodosProvider, useTodos } from './useTodos';
import './App.css';

function TodoInput() {
  const inputRef = useRef();
  const [value, setValue] = useState('');
  const { dispatch } = useTodos();

  const handleAdd = () => {
    const text = inputRef.current.value.trim();
    if (text) {
      dispatch({ type: 'ADD_TODO', text });
      inputRef.current.value = '';
      setValue('');
    }
  };

  return (
    <div style={{ marginBottom: 16 }}>
      <input
        ref={inputRef}
        type="text"
        placeholder="请输入任务"
        defaultValue={value}
        style={{ padding: 8, width: 200, marginRight: 8 }}
      />
      <button onClick={handleAdd} style={{ padding: 8 }}>添加</button>
    </div>
  );
}

function TodoList() {
  const { todos, dispatch } = useTodos();
  const renderedList = useMemo(() => (
    todos.map(todo => (
      <li key={todo.id} style={{ display: 'flex', alignItems: 'center', marginBottom: 8 }}>
        <input
          type="checkbox"
          checked={todo.completed}
          onChange={() => dispatch({ type: 'TOGGLE_TODO', id: todo.id })}
        />
        <span style={{
          textDecoration: todo.completed ? 'line-through' : 'none',
          margin: '0 8px',
          flex: 1
        }}>{todo.text}</span>
        <button onClick={() => dispatch({ type: 'DELETE_TODO', id: todo.id })} style={{ padding: '2px 8px' }}>删除</button>
      </li>
    ))
  ), [todos, dispatch]);

  return <ul style={{ listStyle: 'none', padding: 0 }}>{renderedList}</ul>;
}

function TodoApp() {
  useEffect(() => {
    console.log('Todo List已加载');
  }, []);

  return (
    <TodosProvider>
      <div style={{ maxWidth: 400, margin: '40px auto', padding: 24, background: '#fff', borderRadius: 8, boxShadow: '0 2px 8px rgba(0,0,0,0.1)' }}>
        <h2 style={{ textAlign: 'center' }}>Todo List</h2>
        <TodoInput />
        <TodoList />
      </div>
    </TodosProvider>
  );
}

export default TodoApp;
