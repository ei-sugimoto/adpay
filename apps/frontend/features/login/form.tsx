'use client';

import { loginSchema } from '@/app/login/login.schema';
import { Button } from '@/components/ui/button';
import {
  Form,
  FormControl,
  FormDescription,
  FormField,
  FormItem,
  FormLabel,
  FormMessage,
} from '@/components/ui/form';
import { Input } from '@/components/ui/input';
import { useToast } from '@/hooks/use-toast';
import { nextAPIClient } from '@/lib/apiClient';
import { zodResolver } from '@hookform/resolvers/zod';
import { redirect } from 'next/navigation';
import { useForm } from 'react-hook-form';
import { z } from 'zod';

export default function LoginForm() {
  const form = useForm<z.infer<typeof loginSchema>>({
    resolver: zodResolver(loginSchema),
    defaultValues: {
      name: '',
      password: '',
    },
  });

  const { toast } = useToast();

  async function onSubmit(values: z.infer<typeof loginSchema>) {
    const res = await fetch(`${nextAPIClient.origin}/api/login`, {
      method: 'POST',
      headers: {
        'Content-Type': 'application/json',
      },
      body: JSON.stringify(values),
    }).catch((err) => {
      console.error(err);
      return null;
    });

    if (!res) {
      toast({ title: 'An error occurred', variant: 'destructive' });
      return;
    }

    switch (res.status) {
      case 200:
        redirect('/');
      case 400:
        toast({ title: 'Invalid request', variant: 'destructive' });
        break;
      case 404:
        toast({ title: 'Not found', variant: 'destructive' });
        break;
      default:
        toast({
          title: 'internal server error',
          variant: 'destructive',
        });
        break;
    }
  }

  return (
    <>
      <Form {...form}>
        <form onSubmit={form.handleSubmit(onSubmit)} className='space-y-8'>
          <FormField
            control={form.control}
            name='name'
            render={({ field }) => (
              <FormItem>
                <FormLabel>Username</FormLabel>
                <FormControl>
                  <Input placeholder='user name' {...field} />
                </FormControl>
                <FormDescription>
                  This is your public display name.
                </FormDescription>
                <FormMessage />
              </FormItem>
            )}
          />

          <FormField
            control={form.control}
            name='password'
            render={({ field }) => (
              <FormItem>
                <FormLabel>password</FormLabel>
                <FormControl>
                  <Input placeholder='password' {...field} />
                </FormControl>
                <FormDescription>
                  This is your public display name.
                </FormDescription>
                <FormMessage />
              </FormItem>
            )}
          />

          <Button type='submit'>Submit</Button>
        </form>
      </Form>
    </>
  );
}
