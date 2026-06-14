import React, { useState } from 'react';
import {
  Box,
  Button,
  Card,
  CardContent,
  Container,
  Typography,
  Alert,
  Tabs,
  Tab,
  Grid,
} from '@mui/material';
import {
  Download as DownloadIcon,
  PictureAsPdf as PdfIcon,
  Assessment as AssessmentIcon,
  School as SchoolIcon,
  People as PeopleIcon,
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
      id={`reports-tabpanel-${index}`}
      aria-labelledby={`reports-tab-${index}`}
      {...other}
    >
      {value === index && <Box sx={{ py: 3 }}>{children}</Box>}
    </div>
  );
}

function a11yProps(index: number) {
  return {
    id: `reports-tab-${index}`,
    'aria-controls': `reports-tabpanel-${index}`,
  };
}

const Reports = () => {
  const [tabValue, setTabValue] = useState(0);

  const handleTabChange = (_event: React.SyntheticEvent, newValue: number) => {
    setTabValue(newValue);
  };

  const reportTypes = [
    {
      title: 'School Performance Report',
      description: 'Comprehensive school performance metrics and statistics',
      icon: <SchoolIcon sx={{ fontSize: 48 }} />,
      action: 'Generate Report',
    },
    {
      title: 'Student Achievement Report',
      description: 'Individual and aggregate student achievement data',
      icon: <PeopleIcon sx={{ fontSize: 48 }} />,
      action: 'Generate Report',
    },
    {
      title: 'Assessment Analytics',
      description: 'Detailed analysis of assessment results across all classes',
      icon: <AssessmentIcon sx={{ fontSize: 48 }} />,
      action: 'Generate Report',
    },
    {
      title: 'Curriculum Coverage',
      description: 'Track curriculum completion and coverage across subjects',
      icon: <PdfIcon sx={{ fontSize: 48 }} />,
      action: 'Generate Report',
    },
  ];

  return (
    <Container maxWidth="xl">
      <Box sx={{ mb: 3 }}>
        <Typography variant="h4" component="h1" gutterBottom>
          Reports
        </Typography>
        <Typography variant="body2" color="text.secondary">
          Generate and view various reports for schools, students, and curriculum
        </Typography>
      </Box>

      <Card>
        <CardContent>
          <Box sx={{ borderBottom: 1, borderColor: 'divider' }}>
            <Tabs
              value={tabValue}
              onChange={handleTabChange}
              aria-label="Reports tabs"
            >
              <Tab label="School Reports" {...a11yProps(0)} />
              <Tab label="Student Reports" {...a11yProps(1)} />
              <Tab label="Assessment Reports" {...a11yProps(2)} />
              <Tab label="Curriculum Reports" {...a11yProps(3)} />
            </Tabs>
          </Box>

          <TabPanel value={tabValue} index={0}>
            <Typography variant="h6" gutterBottom>
              School Reports
            </Typography>
            <Grid container spacing={3}>
              {reportTypes.slice(0, 2).map((report, index) => (
                <Grid size={{ xs: 12, md: 6 }} key={index}>
                  <Card variant="outlined">
                    <CardContent>
                      <Box sx={{ display: 'flex', alignItems: 'center', mb: 2 }}>
                        <Box sx={{ mr: 2, color: 'primary.main' }}>
                          {report.icon}
                        </Box>
                        <Box>
                          <Typography variant="h6">{report.title}</Typography>
                          <Typography variant="body2" color="text.secondary">
                            {report.description}
                          </Typography>
                        </Box>
                      </Box>
                      <Button
                        variant="outlined"
                        startIcon={<DownloadIcon />}
                        fullWidth
                      >
                        {report.action}
                      </Button>
                    </CardContent>
                  </Card>
                </Grid>
              ))}
            </Grid>
          </TabPanel>

          <TabPanel value={tabValue} index={1}>
            <Typography variant="h6" gutterBottom>
              Student Reports
            </Typography>
            <Alert severity="info">
              Student reports interface will be implemented here. This will include individual student progress, class-wise performance, and achievement summaries.
            </Alert>
          </TabPanel>

          <TabPanel value={tabValue} index={2}>
            <Typography variant="h6" gutterBottom>
              Assessment Reports
            </Typography>
            <Alert severity="info">
              Assessment reports interface will be implemented here. This will include exam results, assignment scores, and competency progress tracking.
            </Alert>
          </TabPanel>

          <TabPanel value={tabValue} index={3}>
            <Typography variant="h6" gutterBottom>
              Curriculum Reports
            </Typography>
            <Alert severity="info">
              Curriculum reports interface will be implemented here. This will include CP/TP/ATP coverage, modul ajar completion, and curriculum alignment reports.
            </Alert>
          </TabPanel>
        </CardContent>
      </Card>
    </Container>
  );
};

export default Reports;
