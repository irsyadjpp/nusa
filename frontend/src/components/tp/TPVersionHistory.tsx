/**
 * TP Version History Component
 * Displays version history for Teaching Plans
 */

import React, { useState, useEffect } from 'react';
import {
  Box,
  Typography,
  List,
  ListItem,
  ListItemText,
  ListItemIcon,
  Chip,
  Button,
  Dialog,
  DialogTitle,
  DialogContent,
  DialogActions,
  CircularProgress,
  Alert,
} from '@mui/material';
import {
  CheckCircle,
  Info,
} from '@mui/icons-material';
import { getTPSetVersions } from '@/api/tp';
import { TP } from '@/api/tp';

interface TPVersionHistoryProps {
  tpSetId: string;
  open: boolean;
  onClose: () => void;
}

const TPVersionHistory: React.FC<TPVersionHistoryProps> = ({
  tpSetId,
  open,
  onClose,
}) => {
  const [versions, setVersions] = useState<TP[]>([]);
  const [loading, setLoading] = useState(false);
  const [error, setError] = useState<string | null>(null);

  useEffect(() => {
    if (open && tpSetId) {
      loadVersions();
    }
  }, [open, tpSetId]);

  const loadVersions = async () => {
    setLoading(true);
    setError(null);
    try {
      const data = await getTPSetVersions(tpSetId);
      setVersions(data);
    } catch (err: any) {
      setError(err.message || 'Gagal memuat versi TP');
    } finally {
      setLoading(false);
    }
  };

  const getStatusColor = (status: string) => {
    return status === 'APPROVED' ? 'success' : 'default';
  };

  const getStatusIcon = (status: string) => {
    return status === 'APPROVED' ? <CheckCircle color="success" /> : <Info color="action" />;
  };

  return (
    <Dialog open={open} onClose={onClose} maxWidth="md" fullWidth>
      <DialogTitle>Riwayat Versi TP</DialogTitle>
      <DialogContent>
        {loading ? (
          <Box sx={{ display: 'flex', justifyContent: 'center', py: 4 }}>
            <CircularProgress />
          </Box>
        ) : error ? (
          <Alert severity="error" sx={{ mb: 2 }}>
            {error}
          </Alert>
        ) : versions.length === 0 ? (
          <Alert severity="info">
            Tidak ada versi yang ditemukan
          </Alert>
        ) : (
          <List>
            {versions.map((tp) => (
              <ListItem key={tp.id} divider>
                <ListItemIcon>
                  {getStatusIcon(tp.status)}
                </ListItemIcon>
                <ListItemText
                  primary={
                    <Box sx={{ display: 'flex', alignItems: 'center', gap: 2 }}>
                      <Typography variant="subtitle1">
                        {tp.title || 'Tanpa judul'}
                      </Typography>
                      <Chip
                        label={tp.status}
                        color={getStatusColor(tp.status)}
                        size="small"
                      />
                    </Box>
                  }
                  secondary={
                    <Box sx={{ mt: 1 }}>
                      <Typography variant="body2" color="text.secondary">
                        Urutan: {tp.sequence_number}
                      </Typography>
                      <Typography variant="caption" color="text.secondary">
                        Dibuat: {new Date(tp.created_at).toLocaleString('id-ID')}
                      </Typography>
                      {tp.updated_at && (
                        <Typography variant="caption" color="text.secondary">
                          Diupdate: {new Date(tp.updated_at).toLocaleString('id-ID')}
                        </Typography>
                      )}
                    </Box>
                  }
                />
              </ListItem>
            ))}
          </List>
        )}
      </DialogContent>
      <DialogActions>
        <Button onClick={onClose}>Tutup</Button>
      </DialogActions>
    </Dialog>
  );
};

export default TPVersionHistory;
