import React from 'react';
export default function Login() {
  return (
    <div>
      <h2>请先登录</h2>
      <button onClick={() => {
        localStorage.setItem('isLogin', '1');
        window.location.href = '/profile';
      }}>登录</button>
    </div>
  );
}
