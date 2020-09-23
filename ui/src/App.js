import React, { useState } from 'react';
import { Container, Input, Button } from 'semantic-ui-react';

import { showNotification } from './api';

function App() {
  const [notification, setNotification] = useState({ title: '', message: '' });
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
    </Container>
  );
}

export default App;
