/**
 * Class Achievement Page
 * Class-level achievement view
 */

import React, { useState, useEffect } from 'react';
import {
  Box,
  Typography,
  Button,
  Grid,
  CircularProgress,
  Alert,
  TextField,
  MenuItem,
} from '@mui/material';
import {
  ArrowBack,
} from '@mui/icons-material';
import { useNavigate, useParams } from 'react-router-dom';
import { getClassAchievement } from '@/api/achievement';
import ClassAchievementSummary from '@/components/achievement/ClassAchievementSummary';

const ClassAchievementPage: React.FC = () => {
  const navigate = useNavigate();
  const { id } = useParams<{ id: string }>();
  const [classData, setClassData] = useState<any>(null);
  const [loading, setLoading] = useState(true);
  const [error, setError] = useState<string | null>(null);
  const [selectedPeriod, setSelectedPeriod] = useState<string>('current');

  useEffect(() => {
    if (id) {
      loadClassData(id);
    }
  }, [id, selectedPeriod]);

  const loadClassData = async (classId: string) => {
    setLoading(true);
    setError(null);
    try {
      const data = await getClassAchievement(classId, selectedPeriod);
      setClassData(data);
    } catch (err: any) {
      setError(err.message || 'Gagal memuat data kelas');
    } finally {
      setLoading(false);
    }
  };

  if (loading) {
    return (
      <Box sx={{ display: 'flex', justifyContent: 'center', py: 4 }}>
        <CircularProgress />
      </Box>
    );
  }

  if (error) {
    return (
      <Alert severity="error">
        {error}
      </Alert>
    );
  }

  return (
    <Box>
      <Box sx={{ display: 'flex', justifyContent: 'space-between', alignItems: 'center', mb: 3 }}>
        <Box sx={{ display: 'flex', alignItems: 'center', gap: 2 }}>
          <Button
            variant="outlined"
            startIcon={<ArrowBack />}
            onClick={() => navigate('/dashboard')}
          >
            Kembali
          </Button>
          <Typography variant="h4">Pencapaian Kelas</Typography>
        </Box>
        <TextField
          select
          label="Periode"
          value={selectedPeriod}
          onChange={(e) => setSelectedPeriod(e.target.value)}
          sx={{ minWidth: 150 }}
        >
          <MenuItem value="current">Semester Ini</MenuItem>
          <MenuItem value="last">Semester Lalu</MenuItem>
          <MenuItem value="all">Semua</MenuItem>
        </TextField>
      </Box>

      {classData && (
        <ClassAchievementSummary data={classData} />
      )}
    </Box>
  );
};

export default ClassAchievementPage;
