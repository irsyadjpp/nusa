/**
 * Elements List Page
 * Display all curriculum elements with CRUD operations using page-based forms
 */

import React from 'react';
import {
  Box,
  Typography,
  Button,
  Card,
  CardContent,
  CardActions,
  CircularProgress,
  Alert,
  FormControl,
  InputLabel,
  Select,
  MenuItem,
} from '@mui/material';
import {
  Add as AddIcon,
  Delete as DeleteIcon,
} from '@mui/icons-material';
import { useNavigate, useSearchParams } from 'react-router-dom';
import { useElementsByPhase, usePhases } from '@/services/queries/CPQueryService';
import { deleteElement } from '@/api/cp';
import { CurriculumElement } from '@/shared/types/domain';
import { useMutation, useQueryClient } from '@tanstack/react-query';

const ElementsListPage: React.FC = () => {
  const navigate = useNavigate();
  const [searchParams] = useSearchParams();
  const phaseId = searchParams.get('phaseId') || '';
  const queryClient = useQueryClient();

  const handleAdd = () => {
    navigate(`/dashboard/curriculum/elements/new?phaseId=${phaseId}`);
  };

  const handleEdit = (elementId: string) => {
    navigate(`/dashboard/curriculum/elements/${elementId}?phaseId=${phaseId}`);
  };

  const deleteMutation = useMutation({
    mutationFn: deleteElement,
  });

  const handleDelete = async (element: CurriculumElement) => {
    if (window.confirm(`Apakah Anda yakin ingin menghapus elemen "${element.name}"?`)) {
      try {
        await deleteMutation.mutateAsync(element.id);
        queryClient.invalidateQueries({ queryKey: ['elements'] });
      } catch (error) {
        console.error('Error deleting element:', error);
      }
    }
  };

  const { data: elements = [], isLoading, error } = useElementsByPhase(phaseId);
  const { data: phases = [] } = usePhases();

  if (isLoading) {
    return (
      <Box sx={{ display: 'flex', justifyContent: 'center', p: 4 }}>
        <CircularProgress />
      </Box>
    );
  }

  if (error) {
    return <Alert severity="error">Error loading elements</Alert>;
  }

  return (
    <Box sx={{ p: 3 }}>
      <Box sx={{ display: 'flex', justifyContent: 'space-between', alignItems: 'center', mb: 3 }}>
        <Typography variant="h4">Elemen</Typography>
        <Box sx={{ display: 'flex', gap: 2 }}>
          <FormControl sx={{ minWidth: 200 }}>
            <InputLabel>Filter by Phase</InputLabel>
            <Select
              value={phaseId}
              label="Filter by Phase"
              onChange={(e) => {
                navigate(`/dashboard/curriculum/elements?phaseId=${e.target.value}`);
              }}
            >
              <MenuItem value="">All Phases</MenuItem>
              {phases.map((phase) => (
                <MenuItem key={phase.id} value={phase.id}>
                  {phase.name}
                </MenuItem>
              ))}
            </Select>
          </FormControl>
          <Button variant="contained" startIcon={<AddIcon />} onClick={handleAdd}>
            Tambah Elemen
          </Button>
        </Box>
      </Box>

      <Box sx={{ display: 'flex', flexWrap: 'wrap', gap: 3 }}>
        {elements.map((element) => (
          <Box sx={{ width: { xs: '100%', md: '50%', lg: '33.33%' } }} key={element.id}>
            <Card>
              <CardContent>
                <Typography variant="h6" gutterBottom>
                  {element.name}
                </Typography>
                <Typography variant="body2" color="text.secondary" gutterBottom>
                  Kode: {element.code}
                </Typography>
                {element.description && (
                  <Typography variant="body2" color="text.secondary">
                    {element.description}
                  </Typography>
                )}
              </CardContent>
              <CardActions>
                <Button
                  size="small"
                  onClick={() => handleEdit(element.id)}
                >
                  Edit
                </Button>
                <Button
                  size="small"
                  color="error"
                  startIcon={<DeleteIcon />}
                  onClick={() => handleDelete(element)}
                >
                  Hapus
                </Button>
              </CardActions>
            </Card>
          </Box>
        ))}
      </Box>
    </Box>
  );
};

export default ElementsListPage;
