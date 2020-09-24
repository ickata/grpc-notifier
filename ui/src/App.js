import React, { useState, useEffect } from 'react';
import { Container, Input, Button } from 'semantic-ui-react';
import DateTimePicker from 'react-datetime-picker';

import { showNotification, getNotifications, scheduleNotification } from './api';

function App() {
  const [notification, setNotification] = useState({
    title: '',
    message: '',
    datetime: new Date(),
  });
  const [notifications, setNotifications] = useState(null);

  useEffect(() => {
    getNotifications()
      .then(({ notificationsList }) => setNotifications(notificationsList))
      .catch(setNotifications)
  }, []);

  return (
    <Container>
      <h1>Schedule notification</h1>
      <Input
        placeholder="Title"
        onChange={e => setNotification({ ...notification, title: e.target.value })}
      />
      <Input
        placeholder="Message"
        onChange={e => setNotification({ ...notification, message: e.target.value })}
      />
      <Button
        content="Preview"
        onClick={() => showNotification(notification)}
      />
      <DateTimePicker
        value={notification.datetime}
        onChange={datetime => setNotification({ ...notification, datetime })}
      />
      <Button
        primary
        content="Schedule"
        onClick={() => scheduleNotification(notification).then(() => setNotifications([notification, ...notifications]))}
      />

      <h1>Scheduled notifications</h1>
      {!notifications && <h2>Loading...</h2>}
      {notifications instanceof Error && <h2>Error: {notifications.message}</h2>}
      {notifications && !notifications.length && <h2>No notifications yet</h2>}
      {notifications && notifications.length && notifications.map(notification => <div key={notification.title}>
        <h2>{notification.title}</h2>
        <p>{notification.message}</p>
        <p>{new Date(notification.datetime).toString()}</p>
      </div>)}
    </Container>
  );
}

export default App;
