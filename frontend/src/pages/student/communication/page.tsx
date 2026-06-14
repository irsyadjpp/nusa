import React from 'react';
import {
  Box,
  Button,
  Card,
  CardContent,
  Container,
  Typography,
  Chip,
  Tabs,
  Tab,
} from '@mui/material';
import {
  Send as SendIcon,
} from '@mui/icons-material';

interface TabPanelProps {
  children?: React.ReactNode;
  index: number;
  value: number;
}

function TabPanel(props: TabPanelProps) {
  const { children, value, index, ...other } = props;

  return (
    <div
      role="tabpanel"
      hidden={value !== index}
      id={`communication-tabpanel-${index}`}
      aria-labelledby={`communication-tab-${index}`}
      {...other}
    >
      {value === index && <Box sx={{ py: 3 }}>{children}</Box>}
    </div>
  );
}

function a11yProps(index: number) {
  return {
    id: `communication-tab-${index}`,
    'aria-controls': `communication-tabpanel-${index}`,
  };
}

const StudentCommunication = () => {
  const [tabValue, setTabValue] = React.useState(0);

  const handleTabChange = (_event: React.SyntheticEvent, newValue: number) => {
    setTabValue(newValue);
  };

  const notifications = [
    { id: '1', title: 'School Holiday Announcement', date: '2024-01-15', type: 'announcement' },
    { id: '2', title: 'Exam Schedule Update', date: '2024-01-14', type: 'notification' },
    { id: '3', title: 'Library Hours Change', date: '2024-01-13', type: 'announcement' },
  ];

  const messages = [
    { id: '1', from: 'Mr. Smith', subject: 'Mathematics Assignment Feedback', date: '2024-01-15', unread: true },
    { id: '2', from: 'Ms. Johnson', subject: 'Science Lab Report', date: '2024-01-14', unread: false },
    { id: '3', from: 'School Office', subject: 'Fee Payment Reminder', date: '2024-01-13', unread: false },
  ];

  return (
    <Container maxWidth="xl">
      <Box sx={{ mb: 3 }}>
        <Typography variant="h4" component="h1" gutterBottom>
          Communication
        </Typography>
        <Typography variant="body2" color="text.secondary">
          View messages and notifications from teachers and school
        </Typography>
      </Box>

      <Card>
        <CardContent>
          <Box sx={{ borderBottom: 1, borderColor: 'divider' }}>
            <Tabs
              value={tabValue}
              onChange={handleTabChange}
              aria-label="Communication tabs"
            >
              <Tab label="Messages" {...a11yProps(0)} />
              <Tab label="Notifications" {...a11yProps(1)} />
            </Tabs>
          </Box>

          <TabPanel value={tabValue} index={0}>
            <Box sx={{ display: 'flex', justifyContent: 'space-between', alignItems: 'center', mb: 3 }}>
              <Typography variant="h6">Messages</Typography>
              <Button
                variant="contained"
                startIcon={<SendIcon />}
              >
                Compose Message
              </Button>
            </Box>
            {messages.map((message) => (
              <Box
                key={message.id}
                sx={{
                  p: 2,
                  mb: 2,
                  border: 1,
                  borderColor: 'divider',
                  borderRadius: 1,
                  bgcolor: message.unread ? 'action.hover' : 'transparent',
                }}
              >
                <Box sx={{ display: 'flex', justifyContent: 'space-between', alignItems: 'flex-start' }}>
                  <Box>
                    <Typography variant="subtitle1" fontWeight={message.unread ? 'medium' : 'normal'}>
                      {message.subject}
                    </Typography>
                    <Typography variant="body2" color="text.secondary">
                      From: {message.from}
                    </Typography>
                  </Box>
                  <Typography variant="body2" color="text.secondary">
                    {message.date}
                  </Typography>
                </Box>
              </Box>
            ))}
          </TabPanel>

          <TabPanel value={tabValue} index={1}>
            <Box sx={{ display: 'flex', justifyContent: 'space-between', alignItems: 'center', mb: 3 }}>
              <Typography variant="h6">Notifications</Typography>
            </Box>
            {notifications.map((notification) => (
              <Box
                key={notification.id}
                sx={{
                  p: 2,
                  mb: 2,
                  border: 1,
                  borderColor: 'divider',
                  borderRadius: 1,
                }}
              >
                <Box sx={{ display: 'flex', justifyContent: 'space-between', alignItems: 'flex-start' }}>
                  <Box>
                    <Typography variant="subtitle1">{notification.title}</Typography>
                    <Chip
                      label={notification.type}
                      size="small"
                      color={notification.type === 'announcement' ? 'primary' : 'default'}
                      sx={{ mt: 1 }}
                    />
                  </Box>
                  <Typography variant="body2" color="text.secondary">
                    {notification.date}
                  </Typography>
                </Box>
              </Box>
            ))}
          </TabPanel>
        </CardContent>
      </Card>
    </Container>
  );
};

export default StudentCommunication;
