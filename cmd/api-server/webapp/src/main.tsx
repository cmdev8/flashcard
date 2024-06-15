import React from 'react'
import ReactDOM from 'react-dom/client'
import { createBrowserRouter, Navigate, RouterProvider } from "react-router-dom";
import App from './App.tsx'
import './index.css'
import CardsIndex from './pages/cards/CardsIndex.tsx';
import PracticeIndex from './pages/practice/PracticeIndex.tsx';

const router = createBrowserRouter([
  {
    path: "/",
    element: <App />,
    children: [
      {
        path: "/",
        element: <Navigate to="/practice" />,
      },
      {
        path: "/practice",
        element: <PracticeIndex />,
      },
      {
        path: "/cards",
        element: <CardsIndex />,
      },
    ],
  },
  {
    path: "*",
    element: <div>Not found</div>,
  },
]);

ReactDOM.createRoot(document.getElementById('root')!).render(
  <React.StrictMode>
    <RouterProvider router={router} />
  </React.StrictMode>,
)
