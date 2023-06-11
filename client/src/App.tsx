import React from 'react';
import {
  createBrowserRouter,
  RouterProvider,
} from 'react-router-dom';
import './App.css';
import { Header } from './components';
import { DetailedTask, Home, Motivation, News } from './routes';



const router = createBrowserRouter([
  {
    path: "/",
    element: <Home />
  },
  {
    path: "/motivation",
    element: <Motivation />
  },
  {
    path: "/news",
    element: <News />
  },
  {
    path: "/task",
    element: <DetailedTask />
  }
])

function App() {
  return (
    <div className="App">
      <Header/>
      <RouterProvider router={router} />
    </div>
  );
}

export default App;
