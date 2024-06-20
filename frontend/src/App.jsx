import Login from './pages/Login'
import Signup from './pages/SignUp'
import Home from './pages/Home'

import { useEffect, useState } from 'react'
import { BrowserRouter as Router, Routes, Route } from 'react-router-dom'
import axios from 'axios';

const App = () => {
  const [name, setName] = useState('');

  const URL = "http://127.0.0.1:8000/api/getuser"

  useEffect(() => {
    const fetchUser = async () => {
      try {
        const response = await axios.get(URL, {
          headers: {
            'Content-Type': 'application/json'
          },
          credentials: 'include',
        });

        console.log(response);

        const content = await response.json();
        setName(content.name);

      } catch (error) {
        console.error(error)
      }
    }

    fetchUser();

  }, []);


  return (
    <div className='App'>
      <Router>
        <Routes>
          <Route path='/' element={<Home name={name} setName={setName} />} />
          <Route path='/Signup' element={<Signup />} />
          <Route path='/Login' element={<Login setName={setName} />} />
        </Routes>
      </Router>


    </div>
  )
}

export default App;