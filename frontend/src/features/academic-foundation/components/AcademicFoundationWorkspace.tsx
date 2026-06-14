/**
 * Academic Foundation Main Workspace
 * Main component providing tabs and navigation for all academic foundation features
 */

import { useState } from 'react';
import {
  Box,
  Tab,
  Tabs,
  Typography,
  Container,
  Paper,
  Alert,
} from '@mui/material';
import {
  School as SchoolIcon,
  CalendarMonth as CalendarIcon,
  Category as CategoryIcon,
  AccountBalance as ProfileIcon,
  Link as AlignmentIcon,
  Settings as ConfigIcon,
} from '@mui/icons-material';
import { AcademicYearManagement } from './AcademicYearManagement';
import { SemesterManagement } from './SemesterManagement';
import { SubjectCategoryManagement } from './SubjectCategoryManagement';

interface TabPanelProps {
  children?: React.ReactNode;
  index: number;
  value: number;
}

function TabPanel({ children, value, index }: TabPanelProps) {
  return (
    <div role="tabpanel" hidden={value !== index}>
      {value === index && <Box sx={{ py: 3 }}>{children}</Box>}
    </div>
  );
}

export const AcademicFoundationWorkspace = () => {
  const [currentTab, setCurrentTab] = useState(0);
  const [selectedAcademicYear, setSelectedAcademicYear] = useState<string | null>(null);
  const [selectedAcademicYearName, setSelectedAcademicYearName] = useState<string>('');

  const handleTabChange = (_event: React.SyntheticEvent, newValue: number) => {
    setCurrentTab(newValue);
  };

  const handleSelectAcademicYear = (academicYear: any) => {
    setSelectedAcademicYear(academicYear.id);
    setSelectedAcademicYearName(academicYear.name);
    setCurrentTab(1); // Switch to semester tab
  };

  const tabs = [
    {
      label: 'Tahun Ajaran',
      icon: <SchoolIcon />,
      component: (
        <AcademicYearManagement
          schoolId="default-school-id" // TODO: Get from auth context
          onSelectAcademicYear={handleSelectAcademicYear}
        />
      ),
    },
    {
      label: 'Semester',
      icon: <CalendarIcon />,
      component: selectedAcademicYear ? (
        <SemesterManagement
          academicYearId={selectedAcademicYear}
          academicYearName={selectedAcademicYearName}
        />
      ) : (
        <Alert severity="info" sx={{ m: 2 }}>
          Pilih tahun ajaran terlebih dahulu untuk mengelola semester
        </Alert>
      ),
    },
    {
      label: 'Kategori Mata Pelajaran',
      icon: <CategoryIcon />,
      component: <SubjectCategoryManagement />,
    },
    {
      label: 'Profil Lulusan',
      icon: <ProfileIcon />,
      component: (
        <Box sx={{ p: 3 }}>
          <Typography variant="h6" gutterBottom>
            Profil Lulusan
          </Typography>
          <Typography variant="body2" color="text.secondary">
            Kelola dimensi profil lulusan untuk kurikulum
          </Typography>
          <Alert severity="info" sx={{ mt: 2 }}>
            Fitur profil lulusan akan segera tersedia
          </Alert>
        </Box>
      ),
    },
    {
      label: 'CP Alignment',
      icon: <AlignmentIcon />,
      component: (
        <Box sx={{ p: 3 }}>
          <Typography variant="h6" gutterBottom>
            CP Alignment
          </Typography>
          <Typography variant="body2" color="text.secondary">
            Kelola hubungan antara mata pelajaran dan profil lulusan
          </Typography>
          <Alert severity="info" sx={{ mt: 2 }}>
            Fitur CP alignment akan segera tersedia
          </Alert>
        </Box>
      ),
    },
    {
      label: 'Konfigurasi Sistem',
      icon: <ConfigIcon />,
      component: (
        <Box sx={{ p: 3 }}>
          <Typography variant="h6" gutterBottom>
            Konfigurasi Sistem
          </Typography>
          <Typography variant="body2" color="text.secondary">
            Kelola pengaturan sistem global
          </Typography>
          <Alert severity="info" sx={{ mt: 2 }}>
            Fitur konfigurasi sistem akan segera tersedia
          </Alert>
        </Box>
      ),
    },
  ];

  return (
    <Container maxWidth="xl">
      <Box sx={{ mb: 4 }}>
        <Typography variant="h4" component="h1" gutterBottom>
          Academic Foundation
        </Typography>
        <Typography variant="body1" color="text.secondary">
          Kelola struktur akademik dasar: tahun ajaran, semester, mata pelajaran, dan profil lulusan
        </Typography>
      </Box>

      <Paper sx={{ mb: 3 }}>
        <Tabs
          value={currentTab}
          onChange={handleTabChange}
          variant="scrollable"
          scrollButtons="auto"
          sx={{ borderBottom: 1, borderColor: 'divider' }}
        >
          {tabs.map((tab, index) => (
            <Tab
              key={index}
              label={tab.label}
              icon={tab.icon}
              iconPosition="start"
            />
          ))}
        </Tabs>
      </Paper>

      {tabs.map((tab, index) => (
        <TabPanel key={index} value={currentTab} index={index}>
          {tab.component}
        </TabPanel>
      ))}
    </Container>
  );
};