import React from 'react';
import { Outlet, Link } from 'react-router-dom';

export default function Profile() {
  return (
    <div>
      <h2>用户中心</h2>
      <nav>
        <Link to="orders">订单</Link> |{' '}
        <Link to="settings">设置</Link>
      </nav>
      <Outlet />
    </div>
  );
}
