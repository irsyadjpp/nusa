/**
 * Subject Category Management Component
 * Advanced component with tree structure and CRUD operations
 * Now uses page-based forms instead of modals
 */

import { useState } from 'react';
import {
  Box,
  Typography,
  Chip,
  Tooltip,
  IconButton,
  Alert,
  CircularProgress,
  Button,
  Paper,
} from '@mui/material';
import {
  SimpleTreeView,
  TreeItem,
} from '@mui/x-tree-view';
import {
  Add as AddIcon,
  Edit as EditIcon,
  Delete as DeleteIcon,
  Category as CategoryIcon,
} from '@mui/icons-material';
import { useNavigate } from 'react-router-dom';
import { SubjectCategory } from '@/api/academic-foundation';
import { useSubjectCategories } from '@/services/queries/AcademicFoundationQueryService';
import * as academicFoundationApi from '@/api/academic-foundation';

interface SubjectCategoryManagementProps {
  kurikulumVersion?: string;
}

export const SubjectCategoryManagement = ({ kurikulumVersion = '2024' }: SubjectCategoryManagementProps) => {
  const navigate = useNavigate();
  const [expandedItems, setExpandedItems] = useState<string[]>([]);

  const { data: categories, isLoading, error } = useSubjectCategories({
    kurikulum_version: kurikulumVersion
  });

  const handleAddRoot = () => {
    navigate('/dashboard/academic-foundation/subject-categories/new');
  };

  const handleAddChild = (parentId: string) => {
    navigate(`/dashboard/academic-foundation/subject-categories/new?parentId=${parentId}`);
  };

  const handleEdit = (categoryId: string) => {
    navigate(`/dashboard/academic-foundation/subject-categories/${categoryId}`);
  };

  const handleDelete = async (category: SubjectCategory) => {
    if (window.confirm(`Apakah Anda yakin ingin menghapus kategori "${category.name_indonesian}"?`)) {
      try {
        await academicFoundationApi.deleteSubjectCategory(category.id);
        // The page will handle refetching
      } catch (error) {
        console.error('Error deleting category:', error);
      }
    }
  };

  // Build tree structure from flat categories
  const buildTreeStructure = (categories: SubjectCategory[]) => {
    const categoryMap = new Map<string, SubjectCategory & { children: any[] }>();
    
    // Initialize map with empty children arrays
    categories.forEach(category => {
      categoryMap.set(category.id, { ...category, children: [] });
    });

    // Build tree structure
    const tree: any[] = [];
    categories.forEach(category => {
      const categoryNode = categoryMap.get(category.id);
      if (!categoryNode) return;

      if (category.parent_id && categoryMap.has(category.parent_id)) {
        categoryMap.get(category.parent_id)!.children.push(categoryNode);
      } else {
        tree.push(categoryNode);
      }
    });

    return tree;
  };

  const renderTreeItem = (node: any) => {
    const hasChildren = node.children && node.children.length > 0;

    return (
      <TreeItem
        itemId={node.id}
        label={
          <Box sx={{ display: 'flex', alignItems: 'center', justifyContent: 'space-between', pr: 2 }}>
            <Box sx={{ display: 'flex', alignItems: 'center', gap: 1 }}>
              <Typography variant="body2">
                <strong>{node.code}</strong> - {node.name_indonesian}
              </Typography>
              <Chip
                label={`Level ${node.level}`}
                size="small"
                color="info"
              />
              {!node.is_active && (
                <Chip label="Nonaktif" size="small" color="default" />
              )}
            </Box>
            
            <Box sx={{ display: 'flex', gap: 0.5 }}>
              <Tooltip title="Tambah Sub-kategori">
                <IconButton
                  size="small"
                  onClick={(e) => {
                    e.stopPropagation();
                    handleAddChild(node.id);
                  }}
                >
                  <AddIcon fontSize="small" />
                </IconButton>
              </Tooltip>
              
              <Tooltip title="Edit">
                <IconButton
                  size="small"
                  onClick={(e) => {
                    e.stopPropagation();
                    handleEdit(node.id);
                  }}
                >
                  <EditIcon fontSize="small" />
                </IconButton>
              </Tooltip>
              
              <Tooltip title="Hapus">
                <IconButton
                  size="small"
                  color="error"
                  onClick={(e) => {
                    e.stopPropagation();
                    handleDelete(node);
                  }}
                >
                  <DeleteIcon fontSize="small" />
                </IconButton>
              </Tooltip>
            </Box>
          </Box>
        }
      >
        {hasChildren && node.children.map((child: any) => renderTreeItem(child))}
      </TreeItem>
    );
  };

  if (isLoading) {
    return (
      <Box sx={{ display: 'flex', justifyContent: 'center', p: 4 }}>
        <CircularProgress />
      </Box>
    );
  }

  if (error) {
    return (
      <Alert severity="error" sx={{ m: 2 }}>
        Gagal memuat data kategori mata pelajaran
      </Alert>
    );
  }

  const treeData = categories ? buildTreeStructure(categories) : [];

  return (
    <Box sx={{ p: 3 }}>
      <Box sx={{ display: 'flex', justifyContent: 'space-between', alignItems: 'center', mb: 3 }}>
        <Typography variant="h5" component="h2">
          <CategoryIcon sx={{ mr: 1, verticalAlign: 'middle' }} />
          Manajemen Kategori Mata Pelajaran
        </Typography>
        <Button
          variant="contained"
          startIcon={<AddIcon />}
          onClick={handleAddRoot}
        >
          Tambah Kategori Root
        </Button>
      </Box>

      {!categories || categories.length === 0 ? (
        <Box sx={{ textAlign: 'center', py: 8 }}>
          <CategoryIcon sx={{ fontSize: 64, color: 'text.secondary', mb: 2 }} />
          <Typography variant="h6" color="text.secondary" gutterBottom>
            Belum Ada Kategori Mata Pelajaran
          </Typography>
          <Typography variant="body2" color="text.secondary">
            Mulai dengan menambahkan kategori root untuk struktur kurikulum
          </Typography>
        </Box>
      ) : (
        <Paper sx={{ p: 2 }}>
          <SimpleTreeView
            expandedItems={expandedItems}
            onExpandedItemsChange={(_, itemIds) => setExpandedItems(itemIds)}
          >
            {treeData.map((node) => renderTreeItem(node))}
          </SimpleTreeView>
        </Paper>
      )}
    </Box>
  );
};