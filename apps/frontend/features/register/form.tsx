'use client';

import { registerSchema } from '@/app/register/register.schema';
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

export default function RegisterForm() {
  const form = useForm<z.infer<typeof registerSchema>>({
    resolver: zodResolver(registerSchema),
    defaultValues: {
      name: '',
      password: '',
    },
  });

  const { toast } = useToast();

  async function onSubmit(values: z.infer<typeof registerSchema>) {
    await fetch(`${nextAPIClient.origin}/api/register`, {
      method: 'POST',
      headers: {
        'Content-Type': 'application/json',
      },
      body: JSON.stringify(values),
    })
      .catch((err) => {
        console.error(err);
        return null;
      })
      .then((res) => {
        if (!res) {
          return;
        }

        switch (res.status) {
          case 201:
            redirect('/login');
            break;
          case 400:
            toast({ title: 'Invalid request', variant: 'destructive' });
            break;
          case 404:
            toast({ title: 'Not found', variant: 'destructive' });
            break;
          case 409:
            toast({ title: 'already exsists user', variant: 'destructive' });
            break;
          default:
            toast({ title: 'internal server error', variant: 'destructive' });
            break;
        }
      });
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
