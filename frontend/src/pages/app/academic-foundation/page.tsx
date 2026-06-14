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
} from '@mui/material';
import {
  Add as AddIcon,
} from '@mui/icons-material';
import { useNavigate } from 'react-router-dom';

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
      id={`academic-foundation-tabpanel-${index}`}
      aria-labelledby={`academic-foundation-tab-${index}`}
      {...other}
    >
      {value === index && <Box sx={{ py: 3 }}>{children}</Box>}
    </div>
  );
}

function a11yProps(index: number) {
  return {
    id: `academic-foundation-tab-${index}`,
    'aria-controls': `academic-foundation-tabpanel-${index}`,
  };
}

const AcademicFoundation = () => {
  const navigate = useNavigate();
  const [tabValue, setTabValue] = useState(0);

  const handleTabChange = (_event: React.SyntheticEvent, newValue: number) => {
    setTabValue(newValue);
  };

  return (
    <Container maxWidth="xl">
      <Box sx={{ mb: 3 }}>
        <Typography variant="h4" component="h1" gutterBottom>
          Academic Foundation
        </Typography>
        <Typography variant="body2" color="text.secondary">
          Manage curriculum foundation elements: CP, TP, ATP, and Modul Ajar
        </Typography>
      </Box>

      <Card>
        <CardContent>
          <Box sx={{ borderBottom: 1, borderColor: 'divider' }}>
            <Tabs
              value={tabValue}
              onChange={handleTabChange}
              aria-label="Academic Foundation tabs"
            >
              <Tab label="CP (Capaian Pembelajaran)" {...a11yProps(0)} />
              <Tab label="TP (Tujuan Pembelajaran)" {...a11yProps(1)} />
              <Tab label="ATP (Alur Tujuan Pembelajaran)" {...a11yProps(2)} />
              <Tab label="Modul Ajar" {...a11yProps(3)} />
            </Tabs>
          </Box>

          <TabPanel value={tabValue} index={0}>
            <Box sx={{ display: 'flex', justifyContent: 'space-between', alignItems: 'center', mb: 3 }}>
              <Typography variant="h6">Capaian Pembelajaran (Learning Outcomes)</Typography>
              <Button
                variant="contained"
                startIcon={<AddIcon />}
                onClick={() => navigate('/dashboard/academic-foundation/cp/new')}
              >
                Add CP
              </Button>
            </Box>
            <Alert severity="info">
              CP management interface will be implemented here. This will list all Capaian Pembelajaran with filtering and search capabilities.
            </Alert>
          </TabPanel>

          <TabPanel value={tabValue} index={1}>
            <Box sx={{ display: 'flex', justifyContent: 'space-between', alignItems: 'center', mb: 3 }}>
              <Typography variant="h6">Tujuan Pembelajaran (Learning Objectives)</Typography>
              <Button
                variant="contained"
                startIcon={<AddIcon />}
                onClick={() => navigate('/dashboard/academic-foundation/tp/new')}
              >
                Add TP
              </Button>
            </Box>
            <Alert severity="info">
              TP management interface will be implemented here. This will list all Tujuan Pembelajaran with their associated CP and KKTP criteria.
            </Alert>
          </TabPanel>

          <TabPanel value={tabValue} index={2}>
            <Box sx={{ display: 'flex', justifyContent: 'space-between', alignItems: 'center', mb: 3 }}>
              <Typography variant="h6">Alur Tujuan Pembelajaran (Learning Objective Flow)</Typography>
              <Button
                variant="contained"
                startIcon={<AddIcon />}
                onClick={() => navigate('/dashboard/academic-foundation/atp/new')}
              >
                Add ATP
              </Button>
            </Box>
            <Alert severity="info">
              ATP management interface will be implemented here. This will list all Alur Tujuan Pembelajaran showing the sequence of TPs.
            </Alert>
          </TabPanel>

          <TabPanel value={tabValue} index={3}>
            <Box sx={{ display: 'flex', justifyContent: 'space-between', alignItems: 'center', mb: 3 }}>
              <Typography variant="h6">Modul Ajar (Teaching Modules)</Typography>
              <Button
                variant="contained"
                startIcon={<AddIcon />}
                onClick={() => navigate('/dashboard/academic-foundation/modul-ajar/new')}
              >
                Add Modul Ajar
              </Button>
            </Box>
            <Alert severity="info">
              Modul Ajar management interface will be implemented here. This will list all teaching modules with their associated ATPs and assessment components.
            </Alert>
          </TabPanel>
        </CardContent>
      </Card>
    </Container>
  );
};

export default AcademicFoundation;
