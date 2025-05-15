import { useState } from 'react';
import { useNavigate } from 'react-router-dom';

export default function Login() {
  const [isLogin, setIsLogin] = useState(true);
  const [showForgotPopup, setShowForgotPopup] = useState(false);
  const [error, setError] = useState('');
  const navigate = useNavigate();

  const handleSubmit = (e: React.FormEvent) => {
    e.preventDefault();
    setError('');

    const form = e.target as HTMLFormElement;
    const email = (form.elements.namedItem('email') as HTMLInputElement).value;

    // RedHat email check for sign-up
    if (!isLogin && !email.endsWith('@redhat.com')) {
      setError('Only RedHat accounts are currently accepted.');
      return;
    }

    // Simulate success
    navigate('/booking');
  };

  return (
    <div className="min-h-screen bg-[#e5e5e5]">
      {/* HEADER */}
      <header className="bg-white flex items-center px-6 py-4 shadow-sm">
        <div className="flex items-center gap-2 font-bold text-lg">
        <img src="src/images/redhat-logo.png" alt="Red Hat" className="logo-img" />
          <span className="text-black">Red Hat</span>
          <span className="text-black">EV Charging</span>
        </div>
      </header>

      {/* LOGIN BOX */}
      <div className="flex justify-center items-center mt-12 px-4">
        <div className="bg-white w-full max-w-xl rounded-xl shadow-lg flex flex-col justify-between h-[480px] overflow-hidden">
          {/* TOGGLE TABS */}
            <div className="flex bg-gray-400">
            <button
                className={`flex-1 py-3 font-semibold rounded-tl-xl ${
                isLogin ? 'bg-sky-300 text-black' : 'bg-gray-300 text-gray-700'
                }`}
                onClick={() => setIsLogin(true)}
            >
                Log In
            </button>
            <button
                className={`flex-1 py-3 font-semibold rounded-tr-xl ${
                !isLogin ? 'bg-sky-300 text-black' : 'bg-gray-300 text-gray-700'
                }`}
                onClick={() => setIsLogin(false)}
            >
                Sign Up
            </button>
            </div>

          {/* FORM */}
          <form onSubmit={handleSubmit} className="px-10 py-8 space-y-5">
            <div>
              <label className="block text-sm font-medium text-gray-700 mb-1">Email</label>
              <input
                name="email"
                type="email"
                placeholder="Text Input"
                required
                className="w-full px-4 py-2 border border-gray-400 bg-gray-200 rounded"
              />
            </div>
            <div>
              <label className="block text-sm font-medium text-gray-700 mb-1">Password</label>
              <input
                name="password"
                type="password"
                placeholder="Text Input"
                required
                className="w-full px-4 py-2 border border-gray-400 bg-gray-200 rounded"
              />
            </div>
            {error && <p className="text-red-600 text-sm">{error}</p>}

            <button
              type="submit"
              className="w-full bg-red-500 text-white py-2 rounded hover:bg-red-600 transition"
            >
              {isLogin ? 'Log In' : 'Sign Up'}
            </button>

            <div className="text-center text-sm font-semibold text-blue-600 mt-2 space-y-2">
              {isLogin ? (
                <p className="cursor-pointer" onClick={() => setIsLogin(false)}>
                  Register An Account →
                </p>
              ) : (
                <p className="cursor-pointer" onClick={() => setIsLogin(true)}>
                  Already Have An Account? →
                </p>
              )}
              {isLogin && (
                <p
                  className="text-xs underline cursor-pointer"
                  onClick={() => setShowForgotPopup(true)}
                >
                  Forgotten password?
                </p>
              )}
            </div>
          </form>

        {/* FOOTER (always present) */}
          <div className="bg-gray-400 text-center text-sm text-white py-2 rounded-b-xl">
            {isLogin ? <span>&nbsp;</span> : 'Only RedHat accounts currently accepted'}
          </div>
        </div>
      </div>

      {/* FORGOTTEN PASSWORD MODAL */}
      {showForgotPopup && (
        <div className="fixed inset-0 bg-black bg-opacity-40 flex justify-center items-center z-50 px-4">
          <div className="bg-white p-6 rounded-lg shadow-lg max-w-sm w-full">
            <h2 className="text-lg font-semibold mb-4">Forgotten Password</h2>
            <p className="text-sm text-gray-600 mb-4">
              Please contact your RedHat administrator or use internal reset tools.
            </p>
            <button
              className="w-full bg-red-500 text-white py-2 rounded hover:bg-red-600 transition"
              onClick={() => setShowForgotPopup(false)}
            >
              Close
            </button>
          </div>
        </div>
      )}
    </div>
  );
}
