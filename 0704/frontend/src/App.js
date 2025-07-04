import React, { useEffect, useState } from 'react';
import './App.css';

const API_URL = 'http://localhost:8080/books';

function App() {
  const [books, setBooks] = useState([]);
  const [form, setForm] = useState({ id: '', title: '', author: '', stock: '' });
  const [editingId, setEditingId] = useState(null);
  const [loading, setLoading] = useState(false);
  const [error, setError] = useState('');

  // 获取书籍列表
  const fetchBooks = async () => {
    setLoading(true);
    setError('');
    try {
      const res = await fetch(API_URL);
      if (!res.ok) throw new Error('获取书籍失败');
      const data = await res.json();
      setBooks(data);
    } catch (e) {
      setError(e.message);
    } finally {
      setLoading(false);
    }
  };

  useEffect(() => {
    fetchBooks();
  }, []);

  // 表单输入处理
  const handleChange = e => {
    const { name, value } = e.target;
    setForm(f => ({ ...f, [name]: value }));
  };

  // 新增或更新书籍
  const handleSubmit = async e => {
    e.preventDefault();
    setError('');
    if (!form.id || !form.title || !form.author || !form.stock) {
      setError('请填写所有字段');
      return;
    }
    try {
      let res;
      if (editingId) {
        res = await fetch(`${API_URL}/${editingId}`, {
          method: 'PUT',
          headers: { 'Content-Type': 'application/json' },
          body: JSON.stringify(form),
        });
      } else {
        res = await fetch(API_URL, {
          method: 'POST',
          headers: { 'Content-Type': 'application/json' },
          body: JSON.stringify(form),
        });
      }
      if (!res.ok) throw new Error('操作失败');
      setForm({ id: '', title: '', author: '', stock: '' });
      setEditingId(null);
      fetchBooks();
    } catch (e) {
      setError(e.message);
    }
  };

  // 编辑
  const handleEdit = book => {
    setForm(book);
    setEditingId(book.id);
  };

  // 删除
  const handleDelete = async id => {
    if (!window.confirm('确定要删除吗？')) return;
    setError('');
    try {
      const res = await fetch(`${API_URL}/${id}`, { method: 'DELETE' });
      if (!res.ok) throw new Error('删除失败');
      fetchBooks();
    } catch (e) {
      setError(e.message);
    }
  };

  // 取消编辑
  const handleCancel = () => {
    setForm({ id: '', title: '', author: '', stock: '' });
    setEditingId(null);
  };

  return (
    <div className="App">
      <h1>书籍信息管理系统</h1>
      <form className="book-form" onSubmit={handleSubmit}>
        <input name="id" placeholder="ID" value={form.id} onChange={handleChange} disabled={!!editingId} />
        <input name="title" placeholder="书名" value={form.title} onChange={handleChange} />
        <input name="author" placeholder="作者" value={form.author} onChange={handleChange} />
        <input name="stock" placeholder="库存" value={form.stock} onChange={handleChange} />
        <button className="btn" type="submit">{editingId ? '更新' : '新增'}</button>
        {editingId && <button className="btn btn-cancel" type="button" onClick={handleCancel}>取消</button>}
      </form>
      {error && <div className="error">{error}</div>}
      {loading ? <div>加载中...</div> : (
        <table className="book-table">
          <thead>
            <tr>
              <th>ID</th>
              <th>书名</th>
              <th>作者</th>
              <th>库存</th>
              <th>操作</th>
            </tr>
          </thead>
          <tbody>
            {books.map(book => (
              <tr key={book.id}>
                <td>{book.id}</td>
                <td>{book.title}</td>
                <td>{book.author}</td>
                <td>{book.stock}</td>
                <td>
                  <button className="btn btn-edit" onClick={() => handleEdit(book)}>编辑</button>
                  <button className="btn btn-delete" onClick={() => handleDelete(book.id)}>删除</button>
                </td>
              </tr>
            ))}
          </tbody>
        </table>
      )}
    </div>
  );
}

export default App;
