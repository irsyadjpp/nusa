import React, { useState } from 'react';
import { Box, Grid, Paper, Typography, Tabs, Tab, Chip, Button, Stack } from '@mui/material';
import { StatusBadge } from '@/components/shared/StatusBadge';

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
      id={`workflow-tabpanel-${index}`}
      aria-labelledby={`workflow-tab-${index}`}
      {...other}
    >
      {value === index && <Box sx={{ p: 3 }}>{children}</Box>}
    </div>
  );
}

interface PendingItem {
  id: string;
  type: string;
  title: string;
  submittedBy: string;
  submittedAt: string;
  status: string;
}

const WorkflowPage: React.FC = () => {
  const [tabValue, setTabValue] = useState(0);

  const pendingItems: PendingItem[] = [
    {
      id: '1',
      type: 'TP',
      title: 'Mathematics Grade 5 - Teaching Plan',
      submittedBy: 'Teacher A',
      submittedAt: '2024-01-15T10:30:00Z',
      status: 'pending',
    },
    {
      id: '2',
      type: 'Assessment',
      title: 'Science Quiz Chapter 3',
      submittedBy: 'Teacher B',
      submittedAt: '2024-01-15T09:15:00Z',
      status: 'pending',
    },
    {
      id: '3',
      type: 'Modul Ajar',
      title: 'English Module 2',
      submittedBy: 'Teacher C',
      submittedAt: '2024-01-14T16:45:00Z',
      status: 'in_review',
    },
  ];

  const handleApprove = (id: string) => {
    console.log('Approve:', id);
  };

  const handleReject = (id: string) => {
    console.log('Reject:', id);
  };

  const handleView = (id: string) => {
    console.log('View:', id);
  };

  return (
    <Box sx={{ p: 3 }}>
      <Typography variant="h4" component="h1" gutterBottom>
        Workflow - Approval Queue
      </Typography>
      <Typography variant="body1" color="text.secondary" sx={{ mb: 3 }}>
        Review and approve pending submissions
      </Typography>

      <Paper>
        <Tabs
          value={tabValue}
          onChange={(_, newValue) => setTabValue(newValue)}
          aria-label="workflow tabs"
        >
          <Tab label={`Pending Approvals (${pendingItems.length})`} />
          <Tab label="Recently Approved" />
          <Tab label="Recently Rejected" />
        </Tabs>

        <TabPanel value={tabValue} index={0}>
          <Grid container spacing={2}>
            {pendingItems.map((item) => (
              <Grid size={{ xs: 12, md: 6 }} key={item.id}>
                <Paper variant="outlined" sx={{ p: 2 }}>
                  <Box sx={{ display: 'flex', justifyContent: 'space-between', alignItems: 'flex-start', mb: 2 }}>
                    <Box>
                      <Chip label={item.type} size="small" sx={{ mb: 1 }} />
                      <Typography variant="h6">{item.title}</Typography>
                      <Typography variant="body2" color="text.secondary">
                        Submitted by: {item.submittedBy}
                      </Typography>
                      <Typography variant="caption" color="text.secondary">
                        {new Date(item.submittedAt).toLocaleString()}
                      </Typography>
                    </Box>
                    <StatusBadge status={item.status} />
                  </Box>
                  <Stack direction="row" spacing={1} justifyContent="flex-end">
                    <Button size="small" onClick={() => handleView(item.id)}>
                      View
                    </Button>
                    <Button size="small" variant="contained" color="success" onClick={() => handleApprove(item.id)}>
                      Approve
                    </Button>
                    <Button size="small" variant="outlined" color="error" onClick={() => handleReject(item.id)}>
                      Reject
                    </Button>
                  </Stack>
                </Paper>
              </Grid>
            ))}
          </Grid>
        </TabPanel>

        <TabPanel value={tabValue} index={1}>
          <Typography variant="body1" color="text.secondary">
            No recently approved items
          </Typography>
        </TabPanel>

        <TabPanel value={tabValue} index={2}>
          <Typography variant="body1" color="text.secondary">
            No recently rejected items
          </Typography>
        </TabPanel>
      </Paper>
    </Box>
  );
};

export default WorkflowPage;
