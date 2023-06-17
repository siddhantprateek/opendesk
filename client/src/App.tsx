import React, { useContext } from 'react';
import {
  createBrowserRouter,
  RouterProvider,
  Route
} from 'react-router-dom';
import './App.css';
import { Header } from './components';
import { DetailedTask, Home, Motivation, News, Authorize, Dashboard } from './routes';
import { AuthContext } from './context/authprovider';

interface ProtectedRouteProps {
  path: string,
  element: React.ReactElement
}

const ProtectedRoute: React.FC<ProtectedRouteProps> = ({ path, element }) => {
  const { isAuthenticated } = useContext(AuthContext)

  return isAuthenticated ? (
    <Route path={path} element={element} />  
  ) : (
    <Route path='/login' element={<Authorize/>} />
  )
}


const router = createBrowserRouter([
  {
    path: "/",
    element: <Home />
  },
  {
    path: "/login",
    element: <Authorize />
  },
  {
    path: "/motivation",
    element:  <ProtectedRoute path="/motivation" element={<Motivation />} />
  },
  {
    path: "/news",
    element: <ProtectedRoute path="/news" element={<News />} />
  },
  {
    path: "/task",
    element: <ProtectedRoute path="/task" element={<DetailedTask />} />
  },
  {
    path: "/dashboard",
    element: <ProtectedRoute path="/dashboard" element={<Dashboard />} />
  },
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
