import Navbar from "../component/Navbar";

const Home = ({name , setName}) => {
	console.log(name);
	return (
		<>
			<Navbar name={name} setName={setName}/>

			<div className="flex min-h-full w-full text-5xl justify-center items-center ">
				Hello sir , your welcome !
			</div>
		</>
	)
}

export default Home;