import React from 'react';
import { BrowserRouter as Router, Routes, Route, Link } from 'react-router-dom';
import Home from './Home';
import Products from './Products';
import ProductDetail from './ProductDetail';
import Profile from './Profile';
import Orders from './Orders';
import Settings from './Settings';
import Login from './Login';
import NotFound from './NotFound';
import RequireAuth from './RequireAuth';
import './App.css';

function App() {
  return (
    <Router>
      <nav style={{ margin: 10 }}>
        <Link to="/">首页</Link> |{' '}
        <Link to="/products">商品列表</Link> |{' '}
        <Link to="/profile">用户中心</Link>
      </nav>
      <Routes>
        <Route path="/" element={<Home />} />
        <Route path="/products" element={<Products />} />
        <Route path="/products/:id" element={<ProductDetail />} />
        <Route path="/login" element={<Login />} />
        <Route path="/profile" element={
          <RequireAuth>
            <Profile />
          </RequireAuth>
        }>
          <Route path="orders" element={<Orders />} />
          <Route path="settings" element={<Settings />} />
        </Route>
        <Route path="*" element={<NotFound />} />
      </Routes>
    </Router>
  );
}

export default App;
