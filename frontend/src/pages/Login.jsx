import { useState } from "react";
import { auth } from "../firebase";
import { signInWithEmailAndPassword } from "firebase/auth";
import { useNavigate } from "react-router-dom";
import Navbar from "../components/navbar";


function Login() {
  const [email, setEmail] = useState("");
  const [password, setPassword] = useState("");
  const [error, setError] = useState("");
  const navigate = useNavigate();

  const handleLogin = async () => {
    try {
      const userCredential = await signInWithEmailAndPassword(auth, email, password);
      const user = userCredential.user;

      const token = await user.getIdToken();
      console.log("token on the login page : ", token)

      localStorage.setItem("token", token);
      navigate("/dashboard");
    } catch (err) {
      setError(err.message);
    }
  };

  return (
    <div>
      <Navbar showHome={true} />
      <div className="flex justify-center items-center h-screen">
        <div>
          <div>
            <h1 className="bg-blue-400 text-2xl font-bold text-center">🔐 Admin Login</h1>
          </div>
          <div className="bg-amber-200 p-7 max-w-xl">

            <input
              type="email"
              placeholder="Email"
              value={email}
              onChange={(e) => setEmail(e.target.value)}
              className="border p-2 w-full mb-2"
            />
            <input
              type="password"
              placeholder="Password"
              value={password}
              onChange={(e) => setPassword(e.target.value)}
              className="border p-2 w-full mb-2"
            />
            <button
              onClick={handleLogin}
              className="bg-blue-600 text-white px-4 py-2 rounded w-full"
            >
              Login
            </button>
            {error && <p className="text-red-600 mt-2">{error}</p>}
          </div>
        </div>

      </div>
    </div>
  );
}

export default Login;