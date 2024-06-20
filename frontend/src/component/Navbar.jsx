import { Link } from "react-router-dom";
import axios from "axios";

const Navbar = () => {

  const logout = async () => {
    console.log("clicked");
    try {
      const res = await axios.post('http://127.0.0.1:8000/api/logout', {
        method: 'POST',
        headers: { 'Content-Type': 'application/json' },
        withCredentials: true
      });

      console.log(res.data.message);
    } catch (error) {
      console.log(error.response.data.message);
    }

  }


  return (
    <div className="w-full flex justify-between items-center px-20 py-5">
      <div className="">
        <img
          className="mx-auto h-10 w-auto"
          src="https://tailwindui.com/img/logos/mark.svg?color=indigo&shade=600"
          alt="Your Company"
        />
      </div>

      <div className="flex justify-between items-center gap-10">
        <h2>Hello, <b>Devesh sir !</b></h2>
        <Link to="">
          <button
            onClick={() => logout()}
            className="flex  justify-center rounded-md bg-indigo-600 px-3 py-1.5 text-sm font-semibold leading-6 text-white shadow-sm hover:bg-indigo-500 focus-visible:outline focus-visible:outline-2 focus-visible:outline-offset-2 focus-visible:outline-indigo-600"
          >
            Logout
          </button>
        </Link>
      </div>

    </div>
  )
}

export default Navbar