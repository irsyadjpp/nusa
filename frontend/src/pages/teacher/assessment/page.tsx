import React from 'react';
import {
  Box,
  Button,
  Card,
  CardContent,
  Container,
  Typography,
  Grid,
  Chip,
  Tabs,
  Tab,
} from '@mui/material';
import {
  Add as AddIcon,
  Assignment as AssignmentIcon,
  Quiz as QuizIcon,
  School as SchoolIcon,
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
      id={`assessment-tabpanel-${index}`}
      aria-labelledby={`assessment-tab-${index}`}
      {...other}
    >
      {value === index && <Box sx={{ py: 3 }}>{children}</Box>}
    </div>
  );
}

function a11yProps(index: number) {
  return {
    id: `assessment-tab-${index}`,
    'aria-controls': `assessment-tabpanel-${index}`,
  };
}

const TeacherAssessment = () => {
  const [tabValue, setTabValue] = React.useState(0);

  const handleTabChange = (_event: React.SyntheticEvent, newValue: number) => {
    setTabValue(newValue);
  };

  const assessments = [
    {
      id: '1',
      title: 'Mathematics Chapter 5 Quiz',
      subject: 'Mathematics',
      class: '10-A',
      type: 'Quiz',
      date: '2024-01-15',
      status: 'Completed',
      submissions: 30,
      total: 32,
    },
    {
      id: '2',
      title: 'Science Lab Report',
      subject: 'Science',
      class: '10-B',
      type: 'Assignment',
      date: '2024-01-18',
      status: 'In Progress',
      submissions: 20,
      total: 28,
    },
    {
      id: '3',
      title: 'Physics Midterm Exam',
      subject: 'Physics',
      class: '11-A',
      type: 'Exam',
      date: '2024-01-20',
      status: 'Scheduled',
      submissions: 0,
      total: 25,
    },
  ];

  return (
    <Container maxWidth="xl">
      <Box sx={{ mb: 3 }}>
        <Typography variant="h4" component="h1" gutterBottom>
          Assessment Workspace
        </Typography>
        <Typography variant="body2" color="text.secondary">
          Create and manage assessments for your classes
        </Typography>
      </Box>

      <Card>
        <CardContent>
          <Box sx={{ borderBottom: 1, borderColor: 'divider' }}>
            <Tabs
              value={tabValue}
              onChange={handleTabChange}
              aria-label="Assessment tabs"
            >
              <Tab label="My Assessments" {...a11yProps(0)} />
              <Tab label="Create Assessment" {...a11yProps(1)} />
              <Tab label="Grading Queue" {...a11yProps(2)} />
            </Tabs>
          </Box>

          <TabPanel value={tabValue} index={0}>
            <Box sx={{ display: 'flex', justifyContent: 'space-between', alignItems: 'center', mb: 3 }}>
              <Typography variant="h6">My Assessments</Typography>
              <Button
                variant="contained"
                startIcon={<AddIcon />}
                onClick={() => setTabValue(1)}
              >
                Create New Assessment
              </Button>
            </Box>
            <Grid container spacing={3}>
              {assessments.map((assessment) => (
                <Grid size={{ xs: 12, md: 6, lg: 4 }} key={assessment.id}>
                  <Card variant="outlined">
                    <CardContent>
                      <Box sx={{ display: 'flex', justifyContent: 'space-between', alignItems: 'flex-start', mb: 2 }}>
                        <Box sx={{ display: 'flex', alignItems: 'center', gap: 1 }}>
                          {assessment.type === 'Quiz' && <QuizIcon color="primary" />}
                          {assessment.type === 'Assignment' && <AssignmentIcon color="primary" />}
                          {assessment.type === 'Exam' && <SchoolIcon color="primary" />}
                          <Typography variant="subtitle1" fontWeight="medium">
                            {assessment.title}
                          </Typography>
                        </Box>
                        <Chip
                          label={assessment.status}
                          size="small"
                          color={
                            assessment.status === 'Completed'
                              ? 'success'
                              : assessment.status === 'In Progress'
                              ? 'warning'
                              : 'info'
                          }
                        />
                      </Box>
                      <Typography variant="body2" color="text.secondary" gutterBottom>
                        {assessment.subject} • Class {assessment.class}
                      </Typography>
                      <Typography variant="body2" color="text.secondary" gutterBottom>
                        Due: {assessment.date}
                      </Typography>
                      <Box sx={{ display: 'flex', justifyContent: 'space-between', mt: 2 }}>
                        <Typography variant="body2" color="text.secondary">
                          Submissions: {assessment.submissions}/{assessment.total}
                        </Typography>
                        <Button size="small" variant="outlined">
                          View Details
                        </Button>
                      </Box>
                    </CardContent>
                  </Card>
                </Grid>
              ))}
            </Grid>
          </TabPanel>

          <TabPanel value={tabValue} index={1}>
            <Typography variant="h6" gutterBottom>
              Create New Assessment
            </Typography>
            <Box sx={{ maxWidth: 600 }}>
              <Typography variant="body2" color="text.secondary" paragraph>
                Select the type of assessment you want to create:
              </Typography>
              <Grid container spacing={2}>
                <Grid size={{ xs: 12, sm: 4 }}>
                  <Button variant="outlined" fullWidth startIcon={<QuizIcon />}>
                    Quiz
                  </Button>
                </Grid>
                <Grid size={{ xs: 12, sm: 4 }}>
                  <Button variant="outlined" fullWidth startIcon={<AssignmentIcon />}>
                    Assignment
                  </Button>
                </Grid>
                <Grid size={{ xs: 12, sm: 4 }}>
                  <Button variant="outlined" fullWidth startIcon={<SchoolIcon />}>
                    Exam
                  </Button>
                </Grid>
              </Grid>
            </Box>
          </TabPanel>

          <TabPanel value={tabValue} index={2}>
            <Typography variant="h6" gutterBottom>
              Grading Queue
            </Typography>
            <Typography variant="body2" color="text.secondary">
              Assessments pending grading will appear here.
            </Typography>
          </TabPanel>
        </CardContent>
      </Card>
    </Container>
  );
};

export default TeacherAssessment;
