import { useState } from "react";
import axios from "axios";
import Navbar from "../components/navbar";

function Home() {
  const [mode, setMode] = useState("id");
  const [certId, setCertId] = useState("");
  const [form, setForm] = useState({ studentName: "", studentEmail: "", noOfIssue: "", course: "" });
  const [result, setResult] = useState(null);

  const handleSubmit = async () => {
    try {
      const payload = mode === "id"
        ? { certificateHash: certId }
        : form;
      
      const res = await axios.post("http://localhost:8080/verify", payload);
      console.log("Server response:", res.data);
      setResult(res.data);
    } catch (err) {
      console.error("Error during verification:", err);
      setResult("Not Found or Error");
    }
  };

  return (
  <div>
    <Navbar showLogin={true} />
    <div className="flex flex-col items-center justify-center min-h-screen p-4">
      <h1 className="text-3xl font-bold mb-6">🔍 Verify Certificate</h1>

      {/* Mode Switch */}
      <div className="flex gap-4 mb-6">
        <button
          onClick={() => setMode("id")}
          className={`px-4 py-2 rounded ${mode === "id" ? "bg-blue-600 text-white" : "bg-gray-300"}`}
        >
          Verify by ID
        </button>
        <button
          onClick={() => setMode("details")}
          className={`px-4 py-2 rounded ${mode === "details" ? "bg-blue-600 text-white" : "bg-gray-300"}`}
        >
          Verify by Details
        </button>
      </div>

      <div className="bg-amber-200 p-6 rounded-lg shadow-md w-full max-w-md">
        {mode === "id" ? (
          <>
            <h2 className="font-semibold mb-2">Certificate ID</h2>
            <input
              type="text"
              placeholder="Certificate ID"
              value={certId}
              onChange={(e) => setCertId(e.target.value)}
              className="border p-2 w-full mb-4"
            />
          </>
        ) : (
          <>
            <h2 className="font-semibold mb-2">Student Details</h2>
            <input
              type="text"
              placeholder="Student Name"
              value={form.studentName}
              onChange={(e) => setForm({ ...form, studentName: e.target.value })}
              className="border p-2 w-full mb-2"
            />
            <input
              type="email"
              placeholder="Student Email"
              value={form.studentEmail}
              onChange={(e) => setForm({ ...form, studentEmail: e.target.value })}
              className="border p-2 w-full mb-2"
            />
            <input
              type="text"
              placeholder="Issue No"
              value={form.noOfIssue}
              onChange={(e) => setForm({ ...form, noOfIssue: e.target.value })}
              className="border p-2 w-full mb-2"
            />
            <input
              type="text"
              placeholder="Course"
              value={form.course}
              onChange={(e) => setForm({ ...form, course: e.target.value })}
              className="border p-2 w-full mb-4"
            />
          </>
        )}
        
        <button
          onClick={handleSubmit}
          className="bg-blue-600 text-white px-6 py-2 rounded w-full"
        >
          Verify
        </button>
      </div>

      {result && (
        <div className="mt-6 bg-white p-4 shadow-md rounded w-full max-w-md">
          <pre>{JSON.stringify(result, null, 2)}</pre>
        </div>
      )}
    </div>
  </div>
  );
}

export default Home;