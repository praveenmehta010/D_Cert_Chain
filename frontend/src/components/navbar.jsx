import { Link } from "react-router-dom";
import LogoutButton from "./logout";

function Navbar({ showLogin, showLogout, showHome }) {
  return (
    <nav className="bg-blue-600 p-4 flex justify-between items-center">
      <div className="text-white font-bold text-xl">MyApp</div>
      <div className="space-x-4 text-white">
        {showLogin && (
          <div className="flex items-center gap-1.5">
            <p>Need to issue a certificate?</p>
            <Link to="/login" className="hover:underline">Login</Link>
          </div>
        )}
        {showLogout && (
          <div className="flex items-center gap-1.5">
            <p>Gotta logout?</p>
            <LogoutButton />

          </div>
        )}
        {showHome && (
          <div className="flex items-center gap-1.5">
            <p>Go to the home page</p>
            <Link to="/" className="hover:underline">Home</Link>
          </div>
        )}
      </div>
    </nav>
  );
}

export default Navbar;
