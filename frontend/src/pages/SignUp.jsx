
import { useState } from 'react';
import axios from 'axios';
import { useNavigate } from 'react-router-dom';

export default function Signup() {

	const [name, setName] = useState("")
	const [email, setEmail] = useState("")
	const [password, setPassword] = useState("")


	const navigate = useNavigate();


	const SignupHandler = async (e) => {
		e.preventDefault();
		try {
			const Signupdata = { name, email, password }
			const res = await axios.post('http://127.0.0.1:8000/api/register', Signupdata, {
				method: 'POST',
				headers: { 'Content-Type': 'application/json' }
			});
			console.log(res);
			navigate('/login')

		} catch (error) {
			console.log(error);
		}
	}



	return (
		<>

			<div className="flex min-h-full flex-1 flex-col justify-center px-6 py-12 lg:px-8">
				<div className="sm:mx-auto sm:w-full sm:max-w-sm">
					<img
						className="mx-auto h-10 w-auto"
						src="https://tailwindui.com/img/logos/mark.svg?color=indigo&shade=600"
						alt="Your Company"
					/>
					<h2 className="mt-10 text-center text-2xl font-bold leading-9 tracking-tight text-gray-900">
						Create your account
					</h2>
				</div>

				<div className="mt-10 sm:mx-auto sm:w-full sm:max-w-sm">
					<form className="space-y-6" onSubmit={SignupHandler}>

						<div>
							<label htmlFor="name" className="block text-sm font-medium leading-6 text-gray-900">
								FullName:
							</label>
							<div className="mt-2">
								<input
									id="name"
									name="name"
									type="name"
									value={name}
									onChange={(e) => setName(e.target.value)}
									autoComplete="name"
									required
									className="block w-full rounded-md border-0 py-1.5 px-2 text-gray-900 shadow-sm ring-1 ring-inset ring-gray-300 placeholder:text-gray-400 focus:ring-2 focus:ring-inset focus:ring-indigo-600 sm:text-sm sm:leading-6"
								/>
							</div>
						</div>


						<div>
							<label htmlFor="email" className="block text-sm font-medium leading-6 text-gray-900">
								Email address:
							</label>
							<div className="mt-2">
								<input
									id="email"
									name="email"
									type="email"
									value={email}
									onChange={(e) => setEmail(e.target.value)}
									autoComplete="email"
									required
									className="block w-full rounded-md border-0 py-1.5 px-2 text-gray-900 shadow-sm ring-1 ring-inset ring-gray-300 placeholder:text-gray-400 focus:ring-2 focus:ring-inset focus:ring-indigo-600 sm:text-sm sm:leading-6"
								/>
							</div>
						</div>

						<div>
							<div className="flex items-center justify-between">
								<label htmlFor="password" className="block text-sm font-medium leading-6 text-gray-900">
									Password:
								</label>
							</div>
							<div className="mt-2">
								<input
									id="password"
									name="password"
									type="password"
									value={password}
									onChange={(e) => setPassword(e.target.value)}
									autoComplete="current-password"
									required
									className="block w-full rounded-md border-0 py-1.5 px-2  text-gray-900 shadow-sm ring-1 ring-inset ring-gray-300 placeholder:text-gray-400 focus:ring-2 focus:ring-inset focus:ring-indigo-600 sm:text-sm sm:leading-6"
								/>
							</div>
						</div>

						<div>
							<button
								type="submit"
								className="flex w-full justify-center rounded-md bg-indigo-600 px-3 py-1.5 text-sm font-semibold leading-6 text-white shadow-sm hover:bg-indigo-500 focus-visible:outline focus-visible:outline-2 focus-visible:outline-offset-2 focus-visible:outline-indigo-600"
							>
								Sign up
							</button>
						</div>
					</form>


				</div>
			</div>
		</>
	)
}
