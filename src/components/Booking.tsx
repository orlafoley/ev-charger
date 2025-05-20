import React, { useState, useEffect, useRef } from "react";
import { Card, CardContent } from "@/components/ui/card";
import { Button } from "@/components/ui/button";
import { Input } from "@/components/ui/input";
import { CalendarIcon, ClockIcon, LockIcon, PlusIcon } from "lucide-react";
import "./style.css";

export default function Booking() {
  const [date, setDate] = useState("2025-08-17");
  const [time, setTime] = useState("10:30");
  const [duration, setDuration] = useState("60");

  // Notification and user menu state
  const [showNotifications, setShowNotifications] = useState(false);
  const [showUserMenu, setShowUserMenu] = useState(false);

  // Mock user (replace with real auth later)
  const user = { name: "Demo User" };

  // Close dropdowns when clicking outside
  const userMenuRef = useRef(null);
  const notifRef = useRef(null);

  useEffect(() => {
    function handleClickOutside(event: MouseEvent) {
      if (
        userMenuRef.current &&
        !(userMenuRef.current as any).contains(event.target)
      ) {
        setShowUserMenu(false);
      }
      if (
        notifRef.current &&
        !(notifRef.current as any).contains(event.target)
      ) {
        setShowNotifications(false);
      }
    }

    document.addEventListener("mousedown", handleClickOutside);
    return () => document.removeEventListener("mousedown", handleClickOutside);
  }, []);

  // TODO: Make user input dynamic
  const handleQuickBook = async () => {
    try {
      const response = await fetch("http://localhost:8080/api/bookings", {
        method: "POST",
        headers: {
          "Content-Type": "application/json",
        },
        body: JSON.stringify({
          //slot_id: 1, // TODO: Make this dynamic
          name: "Demo User",
          email: "demo@redhat.com",
          date: "2025-08-17",
          time: "11:00",
          duration: 30,
        }),
      });

      if (response.ok) {
        alert("Booking successful!");
      } else {
        alert("Booking failed.");
      }
    } catch (err) {
      console.error("Error booking slot:", err);
      alert("An error occurred.");
    }
  };

  // Regular booking based on form input
  const handleRegularBooking = async () => {
    try {
      const response = await fetch("/api/bookings", {
        method: "POST",
        headers: {
          "Content-Type": "application/json",
        },
        body: JSON.stringify({
          slot_id: 1, // TODO: Make this dynamic
          name: "Demo User",
          email: "demo@redhat.com",
          date,
          time,
          duration: parseInt(duration, 10),
        }),
      });

      if (response.ok) {
        alert("Booking successful!");
      } else {
        alert("Booking failed.");
      }
    } catch (err) {
      console.error("Error booking slot:", err);
      alert("An error occurred.");
    }
  };

  return (
    <div>
      <header>
        <div className="logo">
          <img src="src/images/redhat-logo.png" alt="Red Hat" className="logo-img" />
          Red Hat
        </div>

        {/* TODO: Backend for what happens when you click on these */}
        <nav className="nav">
          <span>
            <svg className="icon" fill="none" stroke="currentColor" strokeWidth="2" viewBox="0 0 24 24" strokeLinecap="round" strokeLinejoin="round">
              <path d="M20 21v-2a4 4 0 0 0-3-3.87" /><path d="M4 21v-2a4 4 0 0 1 3-3.87" /><circle cx="12" cy="7" r="4" /></svg>
            Book
          </span>
          <span>
            <svg className="icon" fill="none" stroke="currentColor" strokeWidth="2" viewBox="0 0 24 24" strokeLinecap="round" strokeLinejoin="round">
              <rect x="3" y="4" width="18" height="18" rx="2" ry="2" /><line x1="16" y1="2" x2="16" y2="6" /><line x1="8" y1="2" x2="8" y2="6" /><line x1="3" y1="10" x2="21" y2="10" /></svg>
            Bookings
          </span>
        </nav>

        <div className="user-info">
          {/* TODO: Logic for notification bell */}
          <button onClick={() => setShowNotifications(!showNotifications)}>
            <BellIcon />
          </button>
          {showNotifications && (
            <div className="notification-dropdown" ref={notifRef}>
              <ul>
                <li>Booking confirmed at 11:00 AM</li>
                <li>Reminder: 30 mins left</li>
                <li>New charger available</li>
              </ul>
            </div>
          )}

          {/* TODO: Logic to click on the account holder's profile */}
          <div className="avatar-wrapper" ref={userMenuRef}>
            <div
              className="avatar"
              onClick={() => setShowUserMenu(!showUserMenu)}
            >
              {user.name.charAt(0).toUpperCase()}
            </div>
            {showUserMenu && (
              <div className="user-dropdown">
                <p>{user.name}</p>
                <button onClick={() => alert("Logging out...")}>Log Out</button>
              </div>
            )}
          </div>
        </div>
      </header>

      <section>
        <h2>Book</h2>
        <div className="grid">
          <div className="input-group">
            <CalendarIcon className="icon" />
            <Input type="date" value={date} onChange={(e) => setDate(e.target.value)} />
          </div>
          <div className="input-group">
            <ClockIcon className="icon" />
            <Input type="time" value={time} onChange={(e) => setTime(e.target.value)} />
          </div>
          <div className="input-group">
            <LockIcon className="icon" />
            <select value={duration} onChange={(e) => setDuration(e.target.value)}>
              <option value="30">30 Mins</option>
              <option value="60">60 Mins</option>
              <option value="90">90 Mins</option>
            </select>
          </div>
        </div>

        {/* Regular Booking Submit Button */}
        <div style={{ marginTop: "1rem" }}>
          <Button className="book-now-button" onClick={handleRegularBooking}>
            Book Now
          </Button>
        </div>
      </section>

      <section className="quick-book">
        <h2>Quick Book</h2>
        <div className="grid">
          {/* TODO: Update this to show different slots rather than the same one three times */}
          {[1, 2, 3].map((item) => (
            <Card key={item} className="card">
              <CardContent>
                <h3>Charger 1</h3>
                <p>17/08/2025</p>
                <p>11:00 AM</p>
                <p>30 Mins</p>
                <div className="button-wrapper">
                  <Button
                    variant="ghost"
                    className="ghost-button"
                    onClick={handleQuickBook}
                  >
                    <PlusIcon />
                  </Button>
                </div>
              </CardContent>
            </Card>
          ))}
        </div>
      </section>
    </div>
  );
}

function BellIcon() {
  return (
    <svg
      width="24"
      height="24"
      fill="none"
      stroke="currentColor"
      strokeWidth="2"
      strokeLinecap="round"
      strokeLinejoin="round"
      className="feather feather-bell"
    >
      <path d="M18 8A6 6 0 0 0 6 8c0 7-3 9-3 9h18s-3-2-3-9" />
      <path d="M13.73 21a2 2 0 0 1-3.46 0" />
    </svg>
  );
}
