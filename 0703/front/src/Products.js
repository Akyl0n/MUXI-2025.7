import React from 'react';
import { Link } from 'react-router-dom';

const products = [
  { id: 1, name: '商品A' },
  { id: 2, name: '商品B' },
  { id: 3, name: '商品C' }
];

export default function Products() {
  return (
    <div>
      <h2>商品列表</h2>
      <ul>
        {products.map(p => (
          <li key={p.id}>
            <Link to={`/products/${p.id}`}>{p.name}</Link>
          </li>
        ))}
      </ul>
    </div>
  );
}
